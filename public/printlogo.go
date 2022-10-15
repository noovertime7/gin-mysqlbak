package public

import "fmt"

var logo = `
  /$$$$$$  /$$                                                               /$$ /$$                /$$      
 /$$__  $$|__/                                                              | $$| $$               | $$      
| $$  \__/ /$$ /$$$$$$$          /$$$$$$/$$$$  /$$   /$$  /$$$$$$$  /$$$$$$ | $$| $$$$$$$  /$$$$$$ | $$   /$$
| $$ /$$$$| $$| $$__  $$ /$$$$$$| $$_  $$_  $$| $$  | $$ /$$_____/ /$$__  $$| $$| $$__  $$|____  $$| $$  /$$/
| $$|_  $$| $$| $$  \ $$|______/| $$ \ $$ \ $$| $$  | $$|  $$$$$$ | $$  \ $$| $$| $$  \ $$ /$$$$$$$| $$$$$$/ 
| $$  \ $$| $$| $$  | $$        | $$ | $$ | $$| $$  | $$ \____  $$| $$  | $$| $$| $$  | $$/$$__  $$| $$_  $$ 
|  $$$$$$/| $$| $$  | $$        | $$ | $$ | $$|  $$$$$$$ /$$$$$$$/|  $$$$$$$| $$| $$$$$$$/  $$$$$$$| $$ \  $$
 \______/ |__/|__/  |__/        |__/ |__/ |__/ \____  $$|_______/  \____  $$|__/|_______/ \_______/|__/  \__/
                                               /$$  | $$                | $$                                 
                                              |  $$$$$$/                | $$                                 
                                               \______/                 |__/                                 `

func PrintLogo() {
	fmt.Println(logo)
	fmt.Println("项目地址: https://github.com/noovertime7/gin-mysqlbak")
	fmt.Println("前端地址: https://github.com/noovertime7/gin-mysqlbak/tree/main/front")
	fmt.Println("集群Client地址: https://github.com/noovertime7/gin-mysqlbak-agent")
}
