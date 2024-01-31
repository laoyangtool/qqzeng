/*
* @Author:  老杨
* @Email:   xcapp1314@gmail.com
* @Date:    2024/1/31 20:34:29 星期三
* @Explain: ...
 */

package qqzeng

type LocationInfo struct {
	Continent   string `json:"continent"`   // 大洲
	CountryZh   string `json:"country"`     // 国家
	CountryEn   string `json:"countryEn"`   // 国家
	CountryCode string `json:"countryCode"` // 国家代码
	Province    string `json:"province"`    // 省份或州
	City        string `json:"city"`        // 城市
	ISP         string `json:"isp"`         // 互联网服务提供商
	Longitude   string `json:"longitude"`   // 经度
	Latitude    string `json:"latitude"`    // 纬度
}
