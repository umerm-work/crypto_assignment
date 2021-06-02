package transform

import "github.com/umerm-work/crypto_assignment/domain"

func ModelToApiPrice(in []domain.CurrencyConversions) domain.Price {

	list := domain.Price{}
	list.RAW = map[string]map[string]domain.Raw{}
	list.DISPLAY = map[string]map[string]domain.Display{}
	mapToCurrencyRaw := map[string]domain.Raw{}
	mapToCurrencyDisplay := map[string]domain.Display{}
	rawCurrencyVal := domain.Raw{}
	displayCurrencyVal := domain.Display{}
	for _, v := range in {
		rawCurrencyVal = domain.Raw{
			CHANGE24HOUR:    v.CHANGEPCT24HOUR,
			CHANGEPCT24HOUR: v.CHANGEPCT24HOUR,
			OPEN24HOUR:      v.OPEN24HOUR,
			VOLUME24HOUR:    v.VOLUME24HOUR,
			VOLUME24HOURTO:  v.VOLUME24HOURTO,
			LOW24HOUR:       v.LOW24HOUR,
			HIGH24HOUR:      v.HIGH24HOUR,
			PRICE:           v.PRICE,
			LASTUPDATE:      v.LASTUPDATE,
			SUPPLY:          v.SUPPLY,
			MKTCAP:          v.MKTCAP,
		}
		displayCurrencyVal = domain.Display{
			CHANGE24HOUR:    v.STRINGCHANGEPCT24HOUR,
			CHANGEPCT24HOUR: v.STRINGCHANGEPCT24HOUR,
			OPEN24HOUR:      v.STRINGOPEN24HOUR,
			VOLUME24HOUR:    v.STRINGVOLUME24HOUR,
			VOLUME24HOURTO:  v.STRINGVOLUME24HOURTO,
			LOW24HOUR:       v.STRINGLOW24HOUR,
			HIGH24HOUR:      v.STRINGHIGH24HOUR,
			PRICE:           v.STRINGPRICE,
			LASTUPDATE:      v.STRINGLASTUPDATE,
			SUPPLY:          v.STRINGSUPPLY,
			MKTCAP:          v.STRINGMKTCAP,
		}
		mapToCurrencyRaw[v.PHYSICALCURRENCYNAME] = rawCurrencyVal
		mapToCurrencyDisplay[v.PHYSICALCURRENCYNAME] = displayCurrencyVal
		list.RAW[v.VIRTUALCURRENCYNAME] = mapToCurrencyRaw
		list.DISPLAY[v.VIRTUALCURRENCYNAME] = mapToCurrencyDisplay
	}
	return list
}

func ApiToModelPrice(in domain.Price) []domain.CurrencyConversions {

	list := []domain.CurrencyConversions{}
	for virtualCurrency, v := range in.RAW {
		for physicalCurrency, timeValues := range v {
			currentVirtualCurr := in.DISPLAY[virtualCurrency]
			list = append(list, domain.CurrencyConversions{
				VIRTUALCURRENCYNAME:  virtualCurrency,
				PHYSICALCURRENCYNAME: physicalCurrency,
				CHANGE24HOUR:         timeValues.CHANGEPCT24HOUR,
				CHANGEPCT24HOUR:      timeValues.CHANGEPCT24HOUR,
				OPEN24HOUR:           timeValues.OPEN24HOUR,
				VOLUME24HOUR:         timeValues.VOLUME24HOUR,
				VOLUME24HOURTO:       timeValues.VOLUME24HOURTO,
				LOW24HOUR:            timeValues.LOW24HOUR,
				HIGH24HOUR:           timeValues.HIGH24HOUR,
				PRICE:                timeValues.PRICE,
				LASTUPDATE:           timeValues.LASTUPDATE,
				SUPPLY:               timeValues.SUPPLY,
				MKTCAP:               timeValues.MKTCAP,

				STRINGCHANGEPCT24HOUR: currentVirtualCurr[physicalCurrency].CHANGEPCT24HOUR,
				STRINGOPEN24HOUR:      currentVirtualCurr[physicalCurrency].CHANGEPCT24HOUR,
				STRINGVOLUME24HOUR:    currentVirtualCurr[physicalCurrency].VOLUME24HOUR,
				STRINGVOLUME24HOURTO:  currentVirtualCurr[physicalCurrency].VOLUME24HOURTO,
				STRINGLOW24HOUR:       currentVirtualCurr[physicalCurrency].LOW24HOUR,
				STRINGHIGH24HOUR:      currentVirtualCurr[physicalCurrency].HIGH24HOUR,
				STRINGPRICE:           currentVirtualCurr[physicalCurrency].PRICE,
				STRINGLASTUPDATE:      currentVirtualCurr[physicalCurrency].LASTUPDATE,
				STRINGSUPPLY:          currentVirtualCurr[physicalCurrency].SUPPLY,
				STRINGMKTCAP:          currentVirtualCurr[physicalCurrency].MKTCAP,
			})
		}
	}
	return list
}
