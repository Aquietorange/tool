package thttp

import (
	"fmt"
	"testing"
)

func Test_New(t *testing.T) {

	/* 	c := New()

	   	ss := c.GetBytes("https://raw.githubusercontent.com/Aquietorange/man2v/master/test/NetPenetrate.sh", "")
	   	fmt.Println(string(ss)) */

	d := New()
	d.header["User-Agent"] = "YoudaoDictPro/9.0.37 (iPhone; iOS 14.8; Scale/3.00)"
	d.header["Content-Type"] = "application/x-www-form-urlencoded"

	/* body := d.PostContent("https://dict.youdao.com/jsonapi?le=eng&keyfrom=voiceGuide&doctype=json&jsonversion=2&dicts=%7b%22count%22%3a36%2c%22dicts%22%3a%5b%5b%22ec%22%5d%2c%5b%22simple%22%5d%5d%7d", "q=juice")
	fmt.Println(body) */

	//individual 默认为中考题型,如没有则为更高级的题型
	body := d.PostContent("https://dict.youdao.com/jsonapi?le=eng&keyfrom=voiceGuide&doctype=json&jsonversion=2&dicts=%7B%22count%22:32,%22dicts%22:%5B%5B%22ec%22%5D,%5B%22ce%22%5D,%5B%22ud%22%5D,%5B%22ec21%22%5D,%5B%22ce_new%22%5D,%5B%22newhh%22%5D,%5B%22collins%22%5D,%5B%22phrs%22%5D,%5B%22rel_word%22%5D,%5B%22pic_dict%22%5D,%5B%22fanyi%22%5D,%5B%22web_search%22%5D,%5B%22typos%22%5D,%5B%22blng_sents_part%22%5D,%5B%22etym%22%5D,%5B%22baike%22%5D,%5B%22exam_dict%22%5D,%5B%22multle%22%5D,%5B%22ugc%22%5D,%5B%22longman%22%5D,%5B%22oxford%22%5D,%5B%22ywAncientWord%22%5D,%5B%22ywBasic%22%5D,%5B%22ywIdiom%22%5D,%5B%22ywRelatedWords%22%5D,%5B%22ywSynAndAnt%22%5D,%5B%22ywWordNet%22%5D,%5B%22newcenturyjc%22%5D,%5B%22newcenturyfc%22%5D,%5B%22word_video%22%5D,%5B%22individual%22%5D%5D%7D", "q=royal")
	fmt.Println(body)

	//mobild 完整请求
	//https://dict.youdao.com/jsonapi_s?client=mobile&le=eng&q=mountain&dicts=%7B%22count%22%3A41%2C%22dicts%22%3A%5B%5B%22ec%22%5D%2C%5B%22ce%22%5D%2C%5B%22ud%22%5D%2C%5B%22ec21%22%5D%2C%5B%22ce_new%22%5D%2C%5B%22ee%22%5D%2C%5B%22newhh%22%5D%2C%5B%22collins%22%5D%2C%5B%22special%22%5D%2C%5B%22phrs%22%5D%2C%5B%22syno%22%5D%2C%5B%22rel_word%22%5D%2C%5B%22pic_dict%22%5D%2C%5B%22fanyi%22%5D%2C%5B%22web_search%22%5D%2C%5B%22typos%22%5D%2C%5B%22web_trans%22%5D%2C%5B%22blng_sents_part%22%5D%2C%5B%22media_sents_part%22%5D%2C%5B%22auth_sents_part%22%5D%2C%5B%22etym%22%5D%2C%5B%22baike%22%5D%2C%5B%22wikipedia_digest%22%5D%2C%5B%22exam_dict%22%5D%2C%5B%22multle%22%5D%2C%5B%22ugc%22%5D%2C%5B%22longman%22%5D%2C%5B%22webster%22%5D%2C%5B%22oxford%22%5D%2C%5B%22oxfordAdvance%22%5D%2C%5B%22ywAncientWord%22%5D%2C%5B%22ywBasic%22%5D%2C%5B%22ywIdiom%22%5D%2C%5B%22ywRelatedWords%22%5D%2C%5B%22ywSynAndAnt%22%5D%2C%5B%22ywWordNet%22%5D%2C%5B%22newcenturyjc%22%5D%2C%5B%22newcenturyfc%22%5D%2C%5B%22video_sents%22%5D%2C%5B%22word_video%22%5D%2C%5B%22individual%22%5D%5D%7D&client=mobile&keyfrom=mdict.9.0.37.iphonepro&imei=4db628be17eb28f2b7f60cb6c4af99d2&model=iPhone&deviceid=4db628be17eb28f2b7f60cb6c4af99d2&mid=14.8&username=18172858836&vendor=AppStore&userid=urs-phoneyd.63d8e5ba5a3049c1b@163.com&device=iPhone_7_Plus&idfa=00000000-0000-0000-0000-000000000000&guestNonce=25969817603678821682&abtest=5&network=wifi&guestSig=5B4B80864B794DCB07590DBA0A4B536C&appVersion=9.0.37&previewEnvTest=0&jsonversion=4&source=main&t=1637118295350&sign=8d40b4ecbb57891f192b0ddbe79e6f90&userLabel=LEARNER,JUNIOR3

}
