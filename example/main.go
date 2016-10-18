package main

import (
	"fmt"

	"github.com/andrewdruzhinin/go-russianpost/russianpost"
	"github.com/jarcoal/httpmock"
)

func main() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://tracking.russianpost.ru/rtm34",
		httpmock.NewStringResponder(200, `
			<?xml version='1.0' encoding='UTF-8'?><S:Envelope xmlns:S="http://www.w3.org/2003/05/soap-envelope"><S:Body><ns7:getOperationHistoryResponse xmlns:ns2="http://russianpost.org/sms-info/data" xmlns:ns3="http://russianpost.org/operationhistory/data" xmlns:ns4="http://schemas.xmlsoap.org/soap/envelope/"
			xmlns:ns5="http://www.russianpost.org/custom-duty-info/data" xmlns:ns6="http://www.russianpost.org/RTM/DataExchangeESPP/Data"
			xmlns:ns7="http://russianpost.org/operationhistory">
			<ns3:OperationHistoryData><ns3:historyRecord>
				<ns3:AddressParameters>
					<ns3:DestinationAddress>
						<ns3:Index>170971</ns3:Index>
						<ns3:Description>Тверь-ДТИ</ns3:Description>
					</ns3:DestinationAddress>
					<ns3:OperationAddress>
						<ns3:Index>443063</ns3:Index>
						<ns3:Description>Самара 63</ns3:Description>
					</ns3:OperationAddress>
					<ns3:MailDirect>
						<ns3:Id>643</ns3:Id>
						<ns3:Code2A>RU</ns3:Code2A>
						<ns3:Code3A>RUS</ns3:Code3A>
						<ns3:NameRU>Российская Федерация</ns3:NameRU>
						<ns3:NameEN>Russian Federation</ns3:NameEN>
					</ns3:MailDirect>
					<ns3:CountryOper>
						<ns3:Id>643</ns3:Id>
						<ns3:Code2A>RU</ns3:Code2A>
						<ns3:Code3A>RUS</ns3:Code3A>
						<ns3:NameRU>Российская Федерация</ns3:NameRU>
						<ns3:NameEN>Russian Federation</ns3:NameEN>
					</ns3:CountryOper>
				</ns3:AddressParameters>
				<ns3:FinanceParameters>
					<ns3:Payment>0</ns3:Payment>
					<ns3:Value>0</ns3:Value>
					<ns3:MassRate>18500</ns3:MassRate>
					<ns3:InsrRate>0</ns3:InsrRate>
					<ns3:AirRate>0</ns3:AirRate>
					<ns3:Rate>1000</ns3:Rate>
					<ns3:CustomDuty>0</ns3:CustomDuty>
				</ns3:FinanceParameters>
				<ns3:ItemParameters>
					<ns3:Barcode>44334455667733</ns3:Barcode>
					<ns3:ValidRuType>false</ns3:ValidRuType>
					<ns3:ValidEnType>false</ns3:ValidEnType>
					<ns3:ComplexItemName>Посылка обыкновенная</ns3:ComplexItemName>
					<ns3:MailRank>
						<ns3:Id>0</ns3:Id>
						<ns3:Name>Без разряда</ns3:Name>
					</ns3:MailRank>
					<ns3:PostMark>
						<ns3:Id>2048</ns3:Id>
						<ns3:Name>Нестандартная</ns3:Name>
					</ns3:PostMark>
					<ns3:MailType>
						<ns3:Id>4</ns3:Id>
						<ns3:Name>Посылка</ns3:Name>
					</ns3:MailType>
					<ns3:MailCtg>
						<ns3:Id>3</ns3:Id>
						<ns3:Name>Обыкновенное</ns3:Name>
					</ns3:MailCtg>
					<ns3:Mass>291</ns3:Mass>
				</ns3:ItemParameters>
				<ns3:OperationParameters>
					<ns3:OperType>
						<ns3:Id>1</ns3:Id>
						<ns3:Name>Прием</ns3:Name>
					</ns3:OperType>
					<ns3:OperAttr>
						<ns3:Id>1</ns3:Id>
						<ns3:Name>Единичный</ns3:Name>
					</ns3:OperAttr>
					<ns3:OperDate>2016-07-16T12:01:00.000+04:00</ns3:OperDate>
				</ns3:OperationParameters>
				<ns3:UserParameters>
					<ns3:SendCtg>
						<ns3:Id>1</ns3:Id>
						<ns3:Name>Население</ns3:Name>
					</ns3:SendCtg>
					<ns3:Sndr></ns3:Sndr>
					<ns3:Rcpn>ООО ДИРЕКТ КАТАЛОГ СЕРВИС</ns3:Rcpn>
				</ns3:UserParameters>
			</ns3:historyRecord>
			<ns3:historyRecord>
				<ns3:AddressParameters>
					<ns3:OperationAddress>
						<ns3:Index>443960</ns3:Index>
						<ns3:Description>Самара МСЦ</ns3:Description>
					</ns3:OperationAddress>
					<ns3:CountryOper>
						<ns3:Id>643</ns3:Id>
						<ns3:Code2A>RU</ns3:Code2A>
						<ns3:Code3A>RUS</ns3:Code3A>
						<ns3:NameRU>Российская Федерация</ns3:NameRU>
						<ns3:NameEN>Russian Federation</ns3:NameEN>
					</ns3:CountryOper>
				</ns3:AddressParameters>
				<ns3:FinanceParameters>
					<ns3:CustomDuty>0</ns3:CustomDuty>
				</ns3:FinanceParameters>
				<ns3:ItemParameters>
					<ns3:Barcode>44334455667733</ns3:Barcode>
					<ns3:ValidRuType>false</ns3:ValidRuType>
					<ns3:ValidEnType>false</ns3:ValidEnType>
				</ns3:ItemParameters>
				<ns3:OperationParameters>
					<ns3:OperType>
						<ns3:Id>8</ns3:Id>
						<ns3:Name>Обработка</ns3:Name>
					</ns3:OperType>
					<ns3:OperAttr>
						<ns3:Id>3</ns3:Id>
						<ns3:Name>Прибыло в сортировочный центр</ns3:Name>
					</ns3:OperAttr>
					<ns3:OperDate>2016-07-18T20:31:27.001+04:00</ns3:OperDate>
				</ns3:OperationParameters>
				<ns3:UserParameters>
					<ns3:Sndr></ns3:Sndr>
					<ns3:Rcpn></ns3:Rcpn>
				</ns3:UserParameters>
			</ns3:historyRecord>
			<ns3:historyRecord>
				<ns3:AddressParameters>
					<ns3:OperationAddress>
						<ns3:Index>443960</ns3:Index>
						<ns3:Description>Самара МСЦ</ns3:Description>
					</ns3:OperationAddress>
					<ns3:CountryOper>
						<ns3:Id>643</ns3:Id>
						<ns3:Code2A>RU</ns3:Code2A>
						<ns3:Code3A>RUS</ns3:Code3A>
						<ns3:NameRU>Российская Федерация</ns3:NameRU>
						<ns3:NameEN>Russian Federation</ns3:NameEN>
					</ns3:CountryOper>
				</ns3:AddressParameters>
				<ns3:FinanceParameters>
					<ns3:CustomDuty>0</ns3:CustomDuty>
				</ns3:FinanceParameters>
				<ns3:ItemParameters>
					<ns3:Barcode>44334455667733</ns3:Barcode>
					<ns3:ValidRuType>false</ns3:ValidRuType>
					<ns3:ValidEnType>false</ns3:ValidEnType>
				</ns3:ItemParameters>
				<ns3:OperationParameters>
					<ns3:OperType>
						<ns3:Id>8</ns3:Id>
						<ns3:Name>Обработка</ns3:Name>
					</ns3:OperType>
					<ns3:OperAttr>
						<ns3:Id>4</ns3:Id>
						<ns3:Name>Покинуло сортировочный центр</ns3:Name>
					</ns3:OperAttr>
					<ns3:OperDate>2016-07-19T06:37:22.001+04:00</ns3:OperDate>
				</ns3:OperationParameters>
				<ns3:UserParameters>
					<ns3:Sndr></ns3:Sndr>
					<ns3:Rcpn></ns3:Rcpn>
				</ns3:UserParameters>
			</ns3:historyRecord>
			<ns3:historyRecord>
				<ns3:AddressParameters>
					<ns3:OperationAddress>
						<ns3:Index>140983</ns3:Index>
						<ns3:Description>Московский АСЦ Цех Посылок</ns3:Description>
					</ns3:OperationAddress>
					<ns3:CountryFrom>
						<ns3:Id>643</ns3:Id>
						<ns3:Code2A>RU</ns3:Code2A>
						<ns3:Code3A>RUS</ns3:Code3A>
						<ns3:NameRU>Российская Федерация</ns3:NameRU>
						<ns3:NameEN>Russian Federation</ns3:NameEN>
					</ns3:CountryFrom>
						<ns3:CountryOper>
						<ns3:Id>643</ns3:Id>
						<ns3:Code2A>RU</ns3:Code2A>
						<ns3:Code3A>RUS</ns3:Code3A>
						<ns3:NameRU>Российская Федерация</ns3:NameRU>
						<ns3:NameEN>Russian Federation</ns3:NameEN>
					</ns3:CountryOper>
				</ns3:AddressParameters>
				<ns3:FinanceParameters>
					<ns3:Payment>0</ns3:Payment>
					<ns3:Value>0</ns3:Value>
					<ns3:MassRate>0</ns3:MassRate>
					<ns3:InsrRate>0</ns3:InsrRate>
					<ns3:AirRate>0</ns3:AirRate>
					<ns3:Rate>0</ns3:Rate>
					<ns3:CustomDuty>0</ns3:CustomDuty>
				</ns3:FinanceParameters>
				<ns3:ItemParameters>
					<ns3:Barcode>44334455667733</ns3:Barcode>
					<ns3:ValidRuType>false</ns3:ValidRuType>
					<ns3:ValidEnType>false</ns3:ValidEnType>
					<ns3:ComplexItemName>Посылка обыкновенная</ns3:ComplexItemName>
					<ns3:MailType>
						<ns3:Id>4</ns3:Id>
						<ns3:Name>Посылка</ns3:Name>
					</ns3:MailType>
					<ns3:MailCtg>
						<ns3:Id>3</ns3:Id>
						<ns3:Name>Обыкновенное</ns3:Name>
					</ns3:MailCtg>
					<ns3:Mass>0</ns3:Mass>
				</ns3:ItemParameters>
				<ns3:OperationParameters>
					<ns3:OperType>
						<ns3:Id>8</ns3:Id>
						<ns3:Name>Обработка</ns3:Name>
					</ns3:OperType>
					<ns3:OperAttr>
						<ns3:Id>0</ns3:Id>
						<ns3:Name>Сортировка</ns3:Name>
					</ns3:OperAttr>
					<ns3:OperDate>2016-07-20T19:33:00.000+03:00</ns3:OperDate>
				</ns3:OperationParameters>
				<ns3:UserParameters>
					<ns3:SendCtg>
						<ns3:Id>0</ns3:Id>
					</ns3:SendCtg>
					<ns3:Sndr></ns3:Sndr>
					<ns3:Rcpn>ООО ДИРЕКТ КАТАЛОГ СЕРВИС</ns3:Rcpn>
				</ns3:UserParameters>
			</ns3:historyRecord>
			<ns3:historyRecord>
				<ns3:AddressParameters>
					<ns3:OperationAddress>
						<ns3:Index>140980</ns3:Index>
						<ns3:Description>Московский АСЦ Цех Логистики</ns3:Description>
					</ns3:OperationAddress>
					<ns3:CountryFrom>
						<ns3:Id>643</ns3:Id>
						<ns3:Code2A>RU</ns3:Code2A>
						<ns3:Code3A>RUS</ns3:Code3A>
						<ns3:NameRU>Российская Федерация</ns3:NameRU>
						<ns3:NameEN>Russian Federation</ns3:NameEN>
					</ns3:CountryFrom>
					<ns3:CountryOper>
						<ns3:Id>643</ns3:Id>
						<ns3:Code2A>RU</ns3:Code2A>
						<ns3:Code3A>RUS</ns3:Code3A>
						<ns3:NameRU>Российская Федерация</ns3:NameRU>
						<ns3:NameEN>Russian Federation</ns3:NameEN>
					</ns3:CountryOper>
					</ns3:AddressParameters>
					<ns3:FinanceParameters>
						<ns3:Payment>0</ns3:Payment>
						<ns3:Value>0</ns3:Value>
						<ns3:MassRate>0</ns3:MassRate>
						<ns3:InsrRate>0</ns3:InsrRate>
						<ns3:AirRate>0</ns3:AirRate>
						<ns3:Rate>0</ns3:Rate>
						<ns3:CustomDuty>0</ns3:CustomDuty>
					</ns3:FinanceParameters>
					<ns3:ItemParameters>
						<ns3:Barcode>44334455667733</ns3:Barcode>
						<ns3:ValidRuType>false</ns3:ValidRuType>
						<ns3:ValidEnType>false</ns3:ValidEnType>
						<ns3:ComplexItemName>Посылка обыкновенная</ns3:ComplexItemName>
						<ns3:MailType>
							<ns3:Id>4</ns3:Id>
							<ns3:Name>Посылка</ns3:Name>
						</ns3:MailType>
						<ns3:MailCtg>
							<ns3:Id>3</ns3:Id>
							<ns3:Name>Обыкновенное</ns3:Name>
						</ns3:MailCtg>
						<ns3:Mass>0</ns3:Mass>
					</ns3:ItemParameters>
					<ns3:OperationParameters>
						<ns3:OperType>
							<ns3:Id>8</ns3:Id>
							<ns3:Name>Обработка</ns3:Name>
						</ns3:OperType>
						<ns3:OperAttr>
							<ns3:Id>4</ns3:Id>
							<ns3:Name>Покинуло сортировочный центр</ns3:Name>
						</ns3:OperAttr>
						<ns3:OperDate>2016-07-21T02:53:00.000+03:00</ns3:OperDate>
					</ns3:OperationParameters>
					<ns3:UserParameters>
						<ns3:SendCtg>
							<ns3:Id>0</ns3:Id>
						</ns3:SendCtg>
						<ns3:Sndr></ns3:Sndr>
						<ns3:Rcpn></ns3:Rcpn>
					</ns3:UserParameters>
				</ns3:historyRecord>
				<ns3:historyRecord>
					<ns3:AddressParameters>
						<ns3:OperationAddress>
							<ns3:Index>170000</ns3:Index>
							<ns3:Description>Тверь Почтамт</ns3:Description>
						</ns3:OperationAddress>
						<ns3:MailDirect>
							<ns3:Id>643</ns3:Id>
							<ns3:Code2A>RU</ns3:Code2A>
							<ns3:Code3A>RUS</ns3:Code3A>
							<ns3:NameRU>Российская Федерация</ns3:NameRU>
							<ns3:NameEN>Russian Federation</ns3:NameEN>
						</ns3:MailDirect>
						<ns3:CountryOper>
							<ns3:Id>643</ns3:Id>
							<ns3:Code2A>RU</ns3:Code2A>
							<ns3:Code3A>RUS</ns3:Code3A>
							<ns3:NameRU>Российская Федерация</ns3:NameRU>
							<ns3:NameEN>Russian Federation</ns3:NameEN>
						</ns3:CountryOper>
					</ns3:AddressParameters>
					<ns3:FinanceParameters>
						<ns3:CustomDuty>0</ns3:CustomDuty>
					</ns3:FinanceParameters>
					<ns3:ItemParameters>
						<ns3:Barcode>44334455667733</ns3:Barcode>
						<ns3:ValidRuType>false</ns3:ValidRuType>
						<ns3:ValidEnType>false</ns3:ValidEnType>
						<ns3:ComplexItemName>Посылка с объявл. ценностью</ns3:ComplexItemName>
						<ns3:MailRank>
							<ns3:Id>0</ns3:Id>
							<ns3:Name>Без разряда</ns3:Name>
						</ns3:MailRank>
						<ns3:MailType>
							<ns3:Id>4</ns3:Id>
							<ns3:Name>Посылка</ns3:Name>
						</ns3:MailType>
						<ns3:MailCtg>
							<ns3:Id>2</ns3:Id>
							<ns3:Name>С объявленной ценностью</ns3:Name>
						</ns3:MailCtg>
						<ns3:Mass>291</ns3:Mass>
					</ns3:ItemParameters>
					<ns3:OperationParameters>
						<ns3:OperType>
							<ns3:Id>8</ns3:Id>
							<ns3:Name>Обработка</ns3:Name>
						</ns3:OperType>
						<ns3:OperAttr>
							<ns3:Id>4</ns3:Id>
							<ns3:Name>Покинуло сортировочный центр</ns3:Name>
						</ns3:OperAttr>
						<ns3:OperDate>2016-07-21T12:24:00.000+03:00</ns3:OperDate>
					</ns3:OperationParameters>
					<ns3:UserParameters>
						<ns3:Sndr></ns3:Sndr>
						<ns3:Rcpn></ns3:Rcpn>
					</ns3:UserParameters>
				</ns3:historyRecord>
				<ns3:historyRecord>
					<ns3:AddressParameters>
						<ns3:OperationAddress>
							<ns3:Index>170100</ns3:Index>
							<ns3:Description>Тверь 100</ns3:Description>
						</ns3:OperationAddress>
						<ns3:MailDirect>
							<ns3:Id>643</ns3:Id>
							<ns3:Code2A>RU</ns3:Code2A>
							<ns3:Code3A>RUS</ns3:Code3A>
							<ns3:NameRU>Российская Федерация</ns3:NameRU>
							<ns3:NameEN>Russian Federation</ns3:NameEN>
						</ns3:MailDirect>
						<ns3:CountryOper>
							<ns3:Id>643</ns3:Id>
							<ns3:Code2A>RU</ns3:Code2A>
							<ns3:Code3A>RUS</ns3:Code3A>
							<ns3:NameRU>Российская Федерация</ns3:NameRU>
							<ns3:NameEN>Russian Federation</ns3:NameEN>
						</ns3:CountryOper>
					</ns3:AddressParameters>
					<ns3:FinanceParameters>
						<ns3:CustomDuty>0</ns3:CustomDuty>
					</ns3:FinanceParameters>
					<ns3:ItemParameters>
						<ns3:Barcode>44334455667733</ns3:Barcode>
						<ns3:ValidRuType>false</ns3:ValidRuType>
						<ns3:ValidEnType>false</ns3:ValidEnType>
						<ns3:ComplexItemName>Посылка с объявл. ценностью</ns3:ComplexItemName>
						<ns3:MailRank>
							<ns3:Id>0</ns3:Id>
							<ns3:Name>Без разряда</ns3:Name>
						</ns3:MailRank>
						<ns3:MailType>
							<ns3:Id>4</ns3:Id>
							<ns3:Name>Посылка</ns3:Name>
						</ns3:MailType>
						<ns3:MailCtg>
							<ns3:Id>2</ns3:Id>
							<ns3:Name>С объявленной ценностью</ns3:Name>
						</ns3:MailCtg>
						<ns3:Mass>291</ns3:Mass>
					</ns3:ItemParameters>
					<ns3:OperationParameters>
						<ns3:OperType>
							<ns3:Id>8</ns3:Id>
							<ns3:Name>Обработка</ns3:Name>
						</ns3:OperType>
						<ns3:OperAttr>
							<ns3:Id>2</ns3:Id>
							<ns3:Name>Прибыло в место вручения</ns3:Name>
						</ns3:OperAttr>
						<ns3:OperDate>2016-07-21T12:39:00.000+03:00</ns3:OperDate>
					</ns3:OperationParameters>
					<ns3:UserParameters>
						<ns3:Sndr></ns3:Sndr>
						<ns3:Rcpn></ns3:Rcpn>
					</ns3:UserParameters>
				</ns3:historyRecord>
				<ns3:historyRecord>
					<ns3:AddressParameters>
						<ns3:OperationAddress>
							<ns3:Index>170100</ns3:Index>
							<ns3:Description>Тверь 100</ns3:Description>
						</ns3:OperationAddress>
						<ns3:MailDirect>
							<ns3:Id>643</ns3:Id>
							<ns3:Code2A>RU</ns3:Code2A>
							<ns3:Code3A>RUS</ns3:Code3A>
							<ns3:NameRU>Российская Федерация</ns3:NameRU>
							<ns3:NameEN>Russian Federation</ns3:NameEN>
						</ns3:MailDirect>
						<ns3:CountryOper>
							<ns3:Id>643</ns3:Id>
							<ns3:Code2A>RU</ns3:Code2A>
							<ns3:Code3A>RUS</ns3:Code3A>
							<ns3:NameRU>Российская Федерация</ns3:NameRU>
							<ns3:NameEN>Russian Federation</ns3:NameEN>
						</ns3:CountryOper>
					</ns3:AddressParameters>
					<ns3:FinanceParameters>
						<ns3:Rate>0</ns3:Rate>
						<ns3:CustomDuty>0</ns3:CustomDuty>
					</ns3:FinanceParameters>
					<ns3:ItemParameters>
						<ns3:Barcode>44334455667733</ns3:Barcode>
						<ns3:ValidRuType>false</ns3:ValidRuType>
						<ns3:ValidEnType>false</ns3:ValidEnType>
						<ns3:ComplexItemName>Посылка с объявл. ценностью</ns3:ComplexItemName>
						<ns3:MailRank>
							<ns3:Id>0</ns3:Id>
							<ns3:Name>Без разряда</ns3:Name>
						</ns3:MailRank>
						<ns3:MailType>
							<ns3:Id>4</ns3:Id>
							<ns3:Name>Посылка</ns3:Name>
						</ns3:MailType>
						<ns3:MailCtg>
							<ns3:Id>2</ns3:Id>
							<ns3:Name>С объявленной ценностью</ns3:Name>
						</ns3:MailCtg>
						<ns3:Mass>291</ns3:Mass>
					</ns3:ItemParameters>
					<ns3:OperationParameters>
						<ns3:OperType>
							<ns3:Id>2</ns3:Id>
							<ns3:Name>Вручение</ns3:Name>
						</ns3:OperType>
						<ns3:OperAttr>
							<ns3:Id>1</ns3:Id>
							<ns3:Name>Вручение адресату</ns3:Name>
							</ns3:OperAttr>
							<ns3:OperDate>2016-07-21T12:49:00.000+03:00</ns3:OperDate>
					</ns3:OperationParameters>
					<ns3:UserParameters>
						<ns3:Sndr></ns3:Sndr>
						<ns3:Rcpn></ns3:Rcpn>
					</ns3:UserParameters>
				</ns3:historyRecord>
			</ns3:OperationHistoryData>
		</ns7:getOperationHistoryResponse>
	</S:Body>
</S:Envelope>`))
	client := russianpost.NewClient("login", "password")
	data, err := client.GetOperationHistory("44334455667733", "0", "RUS")

	if err != nil {
		fmt.Printf("%s", err)
	}
	for _, dataItem := range data.DataItems {
		if dataItem.DestinationAddress != " " {
			fmt.Printf("Пункт назначения: %s \n", dataItem.DestinationAddress)
		}
		fmt.Printf("%s /Вес отправления: %d / %s / %s\n", dataItem.Operation, dataItem.Mass, dataItem.OperationLocation, dataItem.OperarationDate)
	}
}
