package model

import (

	"regexp"
)

type Qimao struct {
	Title string
	Messages string
	Author string
	Point int
}

//获取页面URL
func (q Qimao)QmUrls(html string) []string{
	reg := regexp.MustCompile(`<a.*?href="(/shuku/.*?)"`)
	result :=reg.FindAllStringSubmatch(html,-1)
	var title []string
	for _,val :=range result{
		title = append( title,"https://www.qimao.com"+val[1])
	}
	return title
}
//获取TITLE标题
func(q Qimao) Qmtitle(html string) string  {
	reg := regexp.MustCompile(`<h2 class="tit">(.*?)<`)
	res := reg.FindAllStringSubmatch(html,1)
	return res[0][1]
}

//获取某一部分页面
func(q Qimao) Qmrow(html,math string) [][]string{
	regs := regexp.MustCompile(math)
	res := regs.FindAllStringSubmatch(html,-1)
	return res
}

//获取作者
func(q Qimao) Qmauthor(html string) string{
	reg := regexp.MustCompile(`<em>作者：</em><a.*?>(.*?)<`)
	res := reg.FindAllStringSubmatch(html,-1)
	return res[0][1]
}
//获取主角
func(q Qimao) Qmmast(html string) string{
	reg := regexp.MustCompile(`<em>主角：</em>(.*?)&`)
	res := reg.FindAllStringSubmatch(html,1)
	return res[0][1]
}