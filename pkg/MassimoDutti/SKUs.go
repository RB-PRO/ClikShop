package massimodutti

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Получить список ID товаров по входной ID категории
func SKUs(id_category int) (skus ID, ErrSKU error) {

	url := fmt.Sprintf("https://www.massimodutti.com/itxrest/3/catalog/store/34009471/30359503/category/%v/product?languageId=-1&appId=1&showProducts=false", id_category)

	// fmt.Println("SKUs", url)

	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url, nil)
	if ErrNewRequest != nil {
		return skus, ErrNewRequest
	}
	req.Header.Add("authority", "www.massimodutti.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	// req.Header.Add("cookie", "56311a3c2756d1af186db1478ce849bd=c97d5b19b01d555d28347eb64fd0ed7c; ITXSESSIONID=508974d7d57eba0dc7f2acbbe95b4df8; MDSESSION=136a33ecbb6d01bcb038cb6d339429a3; AKA_A2=A; bm_sz=44D663AC4E355D05D190E8CCAB020F7A~YAAQrWReaMM6WaOJAQAAnhE+pRQmXrv/vXQp+BENAAHAXQFU8YM1Hq/mLtU6GGIBLib9cIvoj/UB5+zohar8qbp8ThnXYjHJP8pqCOh5ejlTCoSe/Q8E8AerDCKq8aJFtuhFPW5qkhYdy0+YU4a0QMYLWzs2Lq/E4bZA+Ek0gHXW2AUwT1F1lnH8bv4+MV6c4wSAuP99UEN3Fb4xbH38rvN9tFQiqrDY0SQweQ6pmSuEAP8Quk71tmS8wUL5wTJXrSoFLOf0nue8FOG90/fH/yGNTnRwx2LKTjCNeioReg5d984+MlGaIpU=~3621938~3420469; bm_mi=5DF82B7384EDC8439E5C629426B20769~YAAQrWReaNE6WaOJAQAAeBI+pRSwj1Z3EXUjefTrxzFgt8AXbU8aTZEojXccemylZ8PgV82KhT9+gm2SP1UNkbVqsQhHxdI2zlYP5BDfu3W/BXWvwfG0HOXaPRbLeZCnYneN20tM6oiamVzMjqfwyS+Vx26gKa28QC2wSBq7jrwryjIl47Y1wZd6RaXD07UEGJi5aikFWteXsrRmtqlMUxTGZtYgxwXAsxNsHd16FKN+uYCtrRMPhhofFve5nM7G8E31hCAIexkFqY6jotm1MMgtFRoQHx7VIu1m4GuA6D3z59p2W2frPPXJlPua14Sb6m9fJksKcjoPgA==~1; JSESSIONID=0000Q-qXGrd595phrEeLoYzbHxh:2aa5bqszw; ak_bmsc=A4BFB4329E0FA4691BF0B16FF482A923~000000000000000000000000000000~YAAQrWReaFQ7WaOJAQAAOB0+pRSMcJvoFw7nrLI236r0mPtHrsE99p028cN4SPOVgCqvXry+rxKpHH7EaTIWtNQPN81ZUECgWw75nMNKgl2//Y8+7BvfOlb/gmUlcntgGjG66rkR+C7mFnXm1Fyr3jU+EI3wZmaXo6Mrj5c1EZC+95Bbzm+o3DMoh2tTUrZ8tqYAIzOPmxNGEhge0MKwcStQGSGswfBVRvTQrGVOnSE88v9qHsoJOWYX3GtB2G3IFjrdrL8yaEsUlHuOuSY6H2R0UEj32wcKT1+A2VHJWy/6aJkrYyE+OQ/IYTwT7zUwjWoqm4zR2r4+5DKW10vxydP2ncfpOZa4uMnL6N0QITbAik6lrzEbdFszmp+3jUx2AU7gtom9JxtYlJIuBiANIy7DP/Xj2D38x5unXUMium5XAtWKvFgtVUur2YLBj98behVtCha/ZwfBB7eR+wJvTbW9upE2Kb8yFcB+cAig1ZNWKuHLSPqOE8JtL/81t0TyCV+viVM+L2jOIi0l7zES+ZLQ6cydng3j; _imutc=8105a29260fabfbfaacf7dac1e9700f94d89be2857eba040c6d414833a42c60e; _gaexp=GAX1.2.E-Z_YWgrSFSF_FBmuRHyuA.19657.4; _gcl_au=1.1.1012536121.1690694461; _gid=GA1.2.812081014.1690694464; FPLC=MgWgd5%2FMxZOTgWRkSmnK6QUL8rz6qC4nWMGsURPDXnt7Sfc4blEoaJaoFTpLagYxbGOLHwf7gOWxsnwgIUUAFX5xGpGDvdohCMDynEPU9fJK2YXYGccMh%2Fq1SQ8vXQ%3D%3D; FPID=FPID2.2.2tLtJwQ%2B%2FIRo9XH3035cWG4NBUzsXnP%2BMzAB5XNSGjM%3D.1690694464; _fbp=fb.1.1690694464596.2064913221; _hjFirstSeen=1; _hjIncludedInSessionSample_1259099=0; _hjSession_1259099=eyJpZCI6Ijc4YzNiZDY4LWMxMWYtNGFmOC1iYzk3LTIxOTU2NDdkYWM2NyIsImNyZWF0ZWQiOjE2OTA2OTQ0NjUzOTUsImluU2FtcGxlIjpmYWxzZX0=; _hjAbsoluteSessionInProgress=0; _pin_unauth=dWlkPVptWmxZV1l4TkRJdE1qaGpOUzAwWldNekxXSTJZemt0Tmprd05XVmlNV0ZsWkRFNQ; _clck=1ehnh5k|2|fdq|0|1306; _abck=617C5E5A38B43DD5163276467C1CCC32~0~YAAQrWReaOxDWaOJAQAACsk+pQqTrRGlAw37GUWZqAqdAoIxEm9sd+MGMUW6MfZ5u7CGaXaMhEnLKwYPTiVDAtXDI/2qlxLX+ukyHbfrnuqI/G7zZ3IpFRTQhzN89HaR/XUNUqPFFfnB+7+wpaZJoZrrYxe3EW8QReYagFMJsLt7Rpt/yOae/zx+nxw9xAtqpN+3wVuAh2HCMOog5P0fLxs1dIHWzSTdWAE1zh9YYna2NtANUyWKHWK5kaoNS+zO4XvP+DzJobb1T4Z+RF7riZrgwXMK+hTDe+nunQX6mC4bwGJpRVpti+OHHSaC6F7XL+hyYXqbxmyZS0OigAX27JBxC0yV9k//NzFKlodWqoC1YR+7YRUK6qWM14yskhANzm+OXmSVxVg9yiEDG8ZDuxQ0t8R7dAAHmxDXR0Y=~-1~-1~-1; _hjSessionUser_1259099=eyJpZCI6IjhhMjI5YzFjLTFjY2MtNWVjMC04ZTFmLWVkYTk4NGVkY2U2OSIsImNyZWF0ZWQiOjE2OTA2OTQ0NjUzNzAsImV4aXN0aW5nIjp0cnVlfQ==; OptanonAlertBoxClosed=2023-07-30T05:45:10.708Z; OptanonConsent=isGpcEnabled=0&datestamp=Sun+Jul+30+2023+08%3A45%3A12+GMT%2B0300+(%D0%9C%D0%BE%D1%81%D0%BA%D0%B2%D0%B0%2C+%D1%81%D1%82%D0%B0%D0%BD%D0%B4%D0%B0%D1%80%D1%82%D0%BD%D0%BE%D0%B5+%D0%B2%D1%80%D0%B5%D0%BC%D1%8F)&version=202211.2.0&isIABGlobal=false&hosts=&consentId=df2ace9f-05f6-4441-8397-6daa6b49cb5f&interactionCount=1&landingPath=NotLandingPage&groups=C0001%3A1%2CC0003%3A1%2CC0002%3A1%2CC0004%3A1%2CC0005%3A1&AwaitingReconsent=false&geolocation=RU%3BMOW; TS01d6af0d=011f37387c08e4d9adabe53491d17d974f74109f3e597c8742cb56c3388d0bdd769c5abd9972179be3d7ac837fe0a1df58365c4921; _ga_F6ZL9NT6KE=GS1.1.1690694464.1.1.1690695939.0.0.0; _ga_NOTFORGA4TRACKING=GS1.1.1690694464.1.1.1690695939.0.0.0; bm_sv=108B0DC15357034C8928F00B5BD9D456~YAAQrWReaEvyWqOJAQAAI7dUpRRq3KEEIC06Fvxf/d5QdiEH5oA03TDMXBC/4SJKfbjb8Y+b9XCTd05WjOhCEsxzmA9FCr5ZzRO2clnRimDkrknbhWOHadRfclPIQVrH+rsSZhwuxM93xiriQHw19ZQO/cojTz0poDUsikuoQ6lCn9XiB//EMAJcAdPGFkunZxfTsuEVyDqquErXHRZRunR9D1XSf1E3SOnwcRy3VMMnxn0H3bbVqPt4QqZmYN53Xe6rsq5BOQ==~1; _uetsid=e13f6b102e9811ee9b9f830fb26b540f; _uetvid=e13ff2c02e9811ee9e98859f58414d5f; _clsk=99j48x|1690695944120|4|0|m.clarity.ms/collect; RT=\"z=1&dm=massimodutti.com&si=0aae6d56-35ff-4d10-a67b-666a1a161a4c&ss=lkoztqrm&sl=6&tt=1vvk&bcn=%2F%2F684dd32e.akstat.io%2F&ld=vyj9&nu=o22kzkh&cl=x545\"; _ga=GA1.2.1525796787.1690694464; _dc_gtm_UA-2867321-1=1; ITXSESSIONID=508974d7d57eba0dc7f2acbbe95b4df8; MDSESSION=136a33ecbb6d01bcb038cb6d339429a3; _abck=617C5E5A38B43DD5163276467C1CCC32~-1~YAAQpGReaJ4YW6OJAQAAKxBjpQpc5+N2XkYBGH42JYbuxJLo3qeBXeZOaSaLqqjXgekhS0PfsNm44fA8Xc7xiMlCwUrR7n1TUmtY5HGDHXrUf7dEwGxFxVBMjCUtX6P+QX/uKCpF/qFQUgoMcI5xjSLZY8hIVqsDuTxCi+s77QW9/id8LHDTmxY3lLpLmQ36nw40MunFxhGbaeBWuSLffScTLzabNQGCT0Frt7feXigdatEkiwQwUAZV5DxkVVGo3VLjk20mAaG5hWxKZeuDhMxZBbFveX52hq2MLIafGXf1FSPvr5rI3ErFbzS61e2kAt0qHwbbh8paSRn0/GzgioaEG5fNWMWFM6inqWjiQvE3X+CspP/w/5MYXWs43igey8AKL2/+tckyIDW5wNB6CrO1JFCJSeGZl8KriXg=~0~-1~-1; bm_sv=108B0DC15357034C8928F00B5BD9D456~YAAQpGReaJ8YW6OJAQAAKxBjpRQOpbMDZ/J79e5isfrIk5yMe5hx0VmKi8hY85k3bxVn0Hbr6CY67bhmjxjUgHe9qf/tLnDl9Tkpzejpeu7GdbFxtygbVExz+E8hJNWtlybIBfDiwcpL9I6nzyFkn+/YSGCMTjVGN15dF/3++seTLhhcmQinF2rYP/M2O34B1u6ECn0dW7D3sD5+/UhkAoY1bxMXpOrsf87sQiPCEn1EVZTbswnRAIGUE6Ho8AcSjqs8UGeX3g==~1; TS01d6af0d=019ceafdc3486a2d99be23eb7f1e048c657a94cd21940a07dec61156c4c70b58e74f44094a8a70371e9d62488eb46a43ffbe6268a4")
	req.Header.Add("referer", "https://www.massimodutti.com/")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.1.906 (beta) Yowser/2.5 Safari/537.36")

	res, ErrDo := client.Do(req)
	if ErrDo != nil {
		return skus, ErrDo
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		ErrNewDecoder := json.NewDecoder(res.Body).Decode(&skus)
		if ErrNewDecoder != nil {
			return skus, ErrNewDecoder
		}
	} else {
		return skus, errors.New("SKUs: http.Status is not ok")
	}
	return skus, nil
}
