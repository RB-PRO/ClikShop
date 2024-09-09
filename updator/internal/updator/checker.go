package updator

import (
	massimodutti "ClikShop/common/MassimoDutti"
	sneaksup "ClikShop/common/SneaSup"
	"ClikShop/common/bases"
	"ClikShop/common/config"
	"ClikShop/common/trendyol"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func (s *Service) CheckParseFromLink(cfg config.Config, link string) (string, error) {
	linkSplit := strings.Split(link, "_")
	if len(linkSplit) != 2 {
		return "", errors.New("error link format " + link)
	}

	number, err := strconv.Atoi(linkSplit[1])
	if err != nil {
		return "", errors.New("error link format " + link)
	}

	switch {
	case strings.Contains(link, bases.TagMD):
		parseLink := cfg.Updater.PingLinks.MassimoDutti[number]
		productLink := strings.ReplaceAll(parseLink, "https://www.massimodutti.com/itxrest/2/catalog/store/34009471/30359503/category/0/product/", "")
		productLink = strings.ReplaceAll(productLink, "/detail?languageId=-1&appId=1", "")
		id, err := strconv.Atoi(productLink)
		if err != nil {
			return "", fmt.Errorf("atoi: %w", err)
		}
		touch, err := s.mdService.Toucher(id)
		if err != nil {
			return "", fmt.Errorf("toucher: %w", err)
		}
		var Product bases.Product2
		Product = massimodutti.Touch2Product2(Product, touch)
		return bases.ProdStr(Product), nil
	case strings.Contains(link, bases.TagHM):
		parseLink := cfg.Updater.PingLinks.HM[number]
		sku := strings.ReplaceAll(parseLink, "https://www2.hm.com/tr_tr/productpage.", "")
		sku = strings.ReplaceAll(sku, ".html", "")
		availability, err := s.hmService.Availability(sku)
		if err != nil {
			return "", fmt.Errorf("hm.Availability: Не получилось получить данные по артикулу %s из ссылки %s: %w", sku, parseLink, err)
		}
		return strings.Join(availability, " "), nil
	case strings.Contains(link, bases.TagZara):
		parseLink := cfg.Updater.PingLinks.Zara[number]
		code := strings.ReplaceAll(parseLink, "https://www.zara.com/tr/en/", "")
		code = strings.ReplaceAll(code, ".html?ajax=true", "")
		product, err := s.zaraService.LoadFantomTouch(code)
		if err != nil {
			return "", fmt.Errorf("touch: %s", err)
		}
		return bases.ProdStr(product), nil
	case strings.Contains(link, bases.TagSS):
		parseLink := cfg.Updater.PingLinks.SneakSup[number]
		colorsItem, err := sneaksup.Aavailability(parseLink)
		if err != nil {
			return "", fmt.Errorf("sneaksup.Availability: %w", err)
		}
		return fmt.Sprintf("%+#v", colorsItem), nil
	case strings.Contains(link, bases.TagTY):
		parseLink := cfg.Updater.PingLinks.Trendyol[number]
		var productID int
		if _, err := fmt.Sscanf(parseLink, trendyol.Product_URL, &productID); err != nil {
			return "", fmt.Errorf("parse link '%s': %w", parseLink, err)
		}
		pg, err := trendyol.ParseProduct(productID)
		if err != nil {
			return "", fmt.Errorf("trendyol.ParseProduct: %w", err)
		}
		return fmt.Sprintf("%+#v", pg), nil
	default:
		return "", fmt.Errorf("не знаю, какую логику применить к тегу %s", link)
	}
}
