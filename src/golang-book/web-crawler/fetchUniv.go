/*
名称：网络爬虫程序
用途：从某大型教育网站抓取大学的信息
作者:腾达格尔
版本：Ver 1.00 Go语言版本：1.3
时间：2014-09-01
*/
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "regexp"
    "time"
)

var (
    tag_BrTag = regexp.MustCompile(`<br>|(</br>)|(<br/>)`)
    tag_PTag  = regexp.MustCompile(`<p>|(</p>)|(&rdquo)|(&ldquo)|(&mdash)|(</strong>)`)

    tag_Nbsp = regexp.MustCompile(`&nbsp;`)

    tag_HTMLTag     = regexp.MustCompile(`(?s)</?.*?>`)
    tag_Space       = regexp.MustCompile(`(^\s+)|( )|(\r\n)|(\t)`)
    tag_SchoolName  = regexp.MustCompile(`\svar\sschoolname='(.+)';\s`)
    tag_SchoolJjUrl = regexp.MustCompile(`<a href="(.+)">学校简介</a>`)
    tag_SchoolJj    = regexp.MustCompile(`<divclass="txt_leftline_24">(.+)</p><scripttype="text/javascript">`)

    //匹配专业列表
    tag_SpecialtyItem = regexp.MustCompile(`<a href="/schoolhtm/specialty/(.+\.htm)" >(.+)</a>`)

    //读取专业内容
    tag_SpecialtyContent = regexp.MustCompile(`<divclass="txt_leftline_24"><p><strong>(.+)</p><scripttype="text/javascript">`)

    //学历层次
    tag_XueLiCengCi = regexp.MustCompile(`<p>学历层次：(.+)<p>招生电话：`)

    //学历层次项目
    tag_XueLiCengCiItem = regexp.MustCompile(`\[(.*?)\]`)

    //招生电话
    tag_Zsdh = regexp.MustCompile(`<p>招生电话：(.+)</p><p>通讯地址：`)

    //通信地址
    tag_Txdz = regexp.MustCompile(`<p>通讯地址：(.+)</p><p>电子邮箱：`)

    //官网地址
    tag_Gwdz = regexp.MustCompile(`<p>官网地址：<ahref="(.+)"target="_blank"class="blue_12">http:`)

    //所在地
    tag_Szd = regexp.MustCompile(`<p>所在地：(.+)院校类型：`)
    //院校类型
    tag_Yxlx = regexp.MustCompile(`院校类型：(.+)</p><p>学历层次：`)
)

//根据网址获得网页内容
func Get(url string) (content string, statusCode int) {
    resp, err1 := http.Get(url)
    if err1 != nil {
        statusCode = -100
        return
    }
    defer resp.Body.Close()
    data, err2 := ioutil.ReadAll(resp.Body)
    if err2 != nil {
        statusCode = -200
        return
    }
    statusCode = resp.StatusCode
    content = string(data)
    return
}

//获得学校首页内容
func GetSchoolHome(code string) (school_home string, ret int) {
    school_url := "http://gkcx.test.com/schoolhtm/schoolTemple/school" + code + ".htm"
    s, statusCode := Get(school_url)
    if statusCode != 200 {
        school_home = ""
        ret = 1
        return
    }
    school_home = s
    ret = 0
    return
}

//获取学校名称
func GetSchoolName(school_info string) (school_name string, xlcc_merge string, zsdh string, txdz string, gwdz string, szd string, yxlx string, ret int) {

    match := tag_SchoolName.FindStringSubmatch(school_info)
    if match != nil {
        school_name = match[1]
        ret = 0
    } else {
        school_name = ""
        ret = 1
    }

    school_info_temp := school_info

    school_info_temp = tag_Space.ReplaceAllString(school_info_temp, "")
    //fmt.Println(school_info_temp)

    //学历层次
    xlcc := ""
    match1 := tag_XueLiCengCi.FindStringSubmatch(school_info_temp)
    if match1 != nil {
        xlcc = match1[1]

    } else {
        xlcc = ""

    }

    matches := tag_XueLiCengCiItem.FindAllStringSubmatch(xlcc, 10000)
    for _, item := range matches {
        xlcc_merge += item[1] + ","

    }
    //fmt.Println("学历层次:",xlcc_merge)

    //招生电话

    matches = tag_Zsdh.FindAllStringSubmatch(school_info_temp, 10000)
    for _, item := range matches {
        zsdh = item[1]

    }
    //fmt.Println("招生电话:",zsdh)

    //所在地
    matches = tag_Szd.FindAllStringSubmatch(school_info_temp, 10000)
    for _, item := range matches {
        szd = item[1]
        os.MkdirAll("data/"+szd, 0666)
    }

    //院校类型
    matches = tag_Yxlx.FindAllStringSubmatch(school_info_temp, 10000)
    for _, item := range matches {
        yxlx = item[1]

    }

    //通信地址

    matches = tag_Txdz.FindAllStringSubmatch(school_info_temp, 10000)
    for _, item := range matches {
        txdz = item[1]

    }
    //fmt.Println("通信地址:",txdz)

    //官网地址

    matches = tag_Gwdz.FindAllStringSubmatch(school_info_temp, 10000)
    for _, item := range matches {
        gwdz = item[1]

    }
    //fmt.Println("官网地址:",gwdz)

    return
}

//获得学校简介
func GetSchoolJj(school_home string) (school_jj string, ret int) {

    match := tag_SchoolJjUrl.FindStringSubmatch(school_home)
    if match != nil {
        school_jj = match[1]
        ret = 0
    } else {
        school_jj = ""
        ret = 1
    }
    school_jj = "http://gkcx.test.com" + school_jj
    //fmt.Println(school_jj)
    s, statusCode := Get(school_jj)
    if statusCode != 200 {
        school_jj = ""
        ret = 1
        return
    }
    school_jj = s
    school_jj = tag_Space.ReplaceAllString(school_jj, "")

    //fmt.Println(school_jj)
    match_jj := tag_SchoolJj.FindStringSubmatch(school_jj)
    if match_jj != nil {
        school_jj = match_jj[1]
        school_jj = tag_PTag.ReplaceAllString(school_jj, "")
        school_jj = tag_BrTag.ReplaceAllString(school_jj, "\r\n")

        ret = 0
    } else {
        school_jj = ""
        ret = 1
    }

    return
}

//专业列表结构体
type SpecialtyItem struct {
    url   string
    title string
}

//获得专业列表
func findSpecialty(code string) (specialty []SpecialtyItem, err error) {

    school_url := "http://gkcx.test.com/schoolhtm/specialty/specialtyList/specialty" + code + ".htm"
    content, statusCode := Get(school_url)
    if statusCode != 200 {
        return
    }

    matches := tag_SpecialtyItem.FindAllStringSubmatch(content, 10000)
    specialty = make([]SpecialtyItem, len(matches))
    for i, item := range matches {
        specialty[i] = SpecialtyItem{"http://gkcx.test.com/schoolhtm/specialty/" + item[1], item[2]}

    }
    return
}

//获得专业内容
func readSpecialty(url string) (content string) {

    content, statusCode := Get(url)
    if statusCode != 200 {
        content = ""
        return
    }
    content = tag_Space.ReplaceAllString(content, "")
    //fmt.Println(content)
    match_jj := tag_SpecialtyContent.FindStringSubmatch(content)
    if match_jj != nil {
        content = match_jj[1]
        content = tag_PTag.ReplaceAllString(content, "")
        content = tag_BrTag.ReplaceAllString(content, "\r\n")
        content = tag_Nbsp.ReplaceAllString(content, " ")

    } else {
        content = ""

    }
    return
}

func GetSchoolInfo(code string) {
    school_home, ret1 := GetSchoolHome(code)
    if ret1 != 0 {
        fmt.Println("Error to get school Home Infomation!", code)
        return
    }

    //获得学校名称
    school_name, xlcc, zsdh, txdz, gwdz, szd, yxlx, ret2 := GetSchoolName(school_home)
    if ret2 != 0 {
        fmt.Println("Error to get school Infomation!", code)
        return
    }

    fmt.Println(szd, school_name, code)
    fileName := fmt.Sprintf("data/%s/列表_%s.txt", szd, school_name)

    //获得学校简介
    school_jj, ret3 := GetSchoolJj(school_home)
    if ret3 != 0 {
        fmt.Println(`Error to GetSchoolJj!`)
        return
    }
    //fmt.Println(school_name,school_jj)

    cnt := fmt.Sprintf("学校名称:%s\n所在地:%s\n院校类型:%s\n学历层次:%s\n招生电话:%s\n通讯地址:%s\n官网地址：%s\n学校简介:\n%s", school_name, szd, yxlx, xlcc, zsdh, txdz, gwdz, school_jj)
    ioutil.WriteFile(fileName, []byte(cnt), 0644)

    //获得专业列表
    specialty, _ := findSpecialty(code)
    for _, item := range specialty {
        //fmt.Printf("获得专业 %s 的内容，来自 %s\n", item.title, item.url)
        specialtyContent := readSpecialty(item.url)
        //fmt.Println(item.title,"====\r\n",specialtyContent)
        fileName := fmt.Sprintf("data/%s/列表_%s_专业_%s.txt", szd, school_name, item.title)
        cnt := fmt.Sprintf("%s 专业： %s \n\n%s", school_name, item.title, specialtyContent)
        ioutil.WriteFile(fileName, []byte(cnt), 0644)
    }
}

//主程序
func main() {
    //根据学校编号遍历所有大学，有些编号可能会不存在
    for i := 30; i < 4000; i++ {
        code := fmt.Sprintf("%d", i)
        if i > 30 && i%50 == 0 {
            fmt.Println("----------------------------")
            time.Sleep(10 * 1e9) //为了减轻网站的负担，每隔50个网站睡眠10秒

        }
        go GetSchoolInfo(code) //并发获得学校的信息
    }

}