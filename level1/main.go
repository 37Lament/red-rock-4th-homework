package main

import "fmt"

func main() {
	var skillNames, xz string
	hp := 100
	for {
		fmt.Printf("圣主在你前面，血量为%d是否使用技能?使用/不使用\n", hp)
		fmt.Scanf("%s", &xz)
		if xz == "不使用" {
			fmt.Println("你落荒而逃")
			break
		} else {
			fmt.Println("您要使用什么技能")
			fmt.Scanf("%s", &skillNames)
			ReleaseSkill(skillNames, func(skillName string) {
				fmt.Println("尝尝我的厉害吧！圣主", skillName)
			})
			hp = hp - 10
			if hp == 0 {
				break

			}
		}
	}
}
func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}
