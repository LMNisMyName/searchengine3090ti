package main

import (
	"fmt"
	"os/exec"
)

func main() {

	command := "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F01169b5823e56ca84a0e282b4e2684.jpg&refer=http%3A%2F%2Fimg.zcool.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1632502166&t=6d560b2d83f33ba20bcde32462850dd1"
	cmd := exec.Command("python3", "img2text.py", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	}
	fmt.Println(string(output))

}
