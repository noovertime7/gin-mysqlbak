package job

import (
	"github.com/noovertime7/mysqlbak/pkg/log"
	"github.com/pkg/errors"
	"sync"
)

type JobType int

var (
	JobFactory  IJobFactory
	factoryOnce sync.Once
)

func GetJobFactory() IJobFactory {
	factoryOnce.Do(func() {
		JobFactory = newJobFactory()
	})
	return JobFactory
}

type IJobFactory interface {
	Start() error
	Stop(JobType) error
	GetJobber(JobType) (Jobber, error)
	Register(Jobber, JobType)
	IsSuccess(JobType) bool
}

func newJobFactory() IJobFactory {
	return &jobFactory{
		lock:      sync.Mutex{},
		jobs:      make(map[JobType]Jobber),
		startJobs: make(map[JobType]bool),
	}
}

type jobFactory struct {
	lock          sync.Mutex
	defaultReSync string
	jobs          map[JobType]Jobber
	startJobs     map[JobType]bool
}

func (j *jobFactory) IsSuccess(jType JobType) bool {
	j.lock.Lock()
	defer j.lock.Unlock()
	errs, ok := j.jobs[jType].GetErr()
	for _, e := range errs {
		log.Logger.Errorf("%s Job err %s", jType, e)
	}
	return ok
}

func (j *jobFactory) Start() error {
	j.lock.Lock()
	defer j.lock.Unlock()
	if len(j.jobs) == 0 {
		return errors.New("当前没有任何服务注册")
	}
	for jobType, job := range j.jobs {
		if !j.startJobs[jobType] {
			go job.Start()
			j.startJobs[jobType] = true
		}
	}
	return nil
}

func (j *jobFactory) Register(job Jobber, jobType JobType) {
	j.lock.Lock()
	defer j.lock.Unlock()
	j.jobs[jobType] = job
}

func (j *jobFactory) Stop(jobType JobType) error {
	j.lock.Lock()
	defer j.lock.Unlock()
	if !j.isExists(jobType) {
		return errors.New("任务未注册")
	}
	for jType, job := range j.jobs {
		if jType == jobType {
			job.Stop()
			j.startJobs[jobType] = false
		}
	}
	return nil
}

func (j *jobFactory) GetJobber(jType JobType) (Jobber, error) {
	j.lock.Lock()
	defer j.lock.Unlock()
	if !j.isExists(jType) {
		return nil, errors.New("任务未注册")
	}
	return j.jobs[jType], nil
}

func (j *jobFactory) isExists(jobType JobType) bool {
	_, ok := j.jobs[jobType]
	return ok
}
