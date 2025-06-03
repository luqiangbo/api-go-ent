package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("开始生成Ent代码...")

	cmd := exec.Command("go", "generate", "./ent")
	cmd.Dir = "../../" // 回到项目根目录
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("生成Ent代码失败: %v", err)
	}

	fmt.Println("Ent代码生成完成")
}
