/*
* CODE GENERATED AUTOMATICALLY WITH github.com/gregpechiro/neo4jGenerator
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
)

var NoCreditCardFound = fmt.Errorf("no creditCard found")
var MultipleCreditCardFound = fmt.Errorf("multiple creditCard found")

func AddCreditCard(creditCard CreditCard) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE (creditCard:CreditCard { Id:{creditCardId}, Number:{creditCardNumber}, ExpirationDate:{creditCardExpirationDate}, SecurityCode:{creditCardSecurityCode}, NameOnCard:{creditCardNameOnCard}, Typ:{creditCardTyp} })", map[string]interface{}{
		"creditCardId":             creditCard.Id,
		"creditCardNumber":         creditCard.Number,
		"creditCardExpirationDate": creditCard.ExpirationDate,
		"creditCardSecurityCode":   creditCard.SecurityCode,
		"creditCardNameOnCard":     creditCard.NameOnCard,
		"creditCardTyp":            creditCard.Typ,
	})

	return err
}

func GetAllCreditCard() ([]CreditCard, error) {
	var creditCards []CreditCard
	conn, err := driver.OpenPool()
	if err != nil {
		return creditCards, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard) RETURN creditCard", nil)
	if err != nil {
		return creditCards, err
	}
	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return creditCards, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return creditCards, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		creditCard := CreditCard{}
		if Id, ok := node.Properties["Id"].(string); ok {
			creditCard.Id = Id
		}
		if Number, ok := node.Properties["Number"].(string); ok {
			creditCard.Number = Number
		}
		if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
			creditCard.ExpirationDate = ExpirationDate
		}
		if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
			creditCard.SecurityCode = SecurityCode
		}
		if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
			creditCard.NameOnCard = NameOnCard
		}
		if Typ, ok := node.Properties["Typ"].(string); ok {
			creditCard.Typ = Typ
		}

		creditCards = append(creditCards, creditCard)
	}

	return creditCards, nil
}

func GetCreditCardById(id string) (CreditCard, error) {
	creditCard := CreditCard{}

	conn, err := driver.OpenPool()
	if err != nil {
		return creditCard, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ Id:{ Id } }) RETURN CreditCard", map[string]interface{}{
		"Id": id,
	})
	if err != nil {
		return creditCard, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return creditCard, NoCreditCardFound
	}
	if err != nil {
		return creditCard, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return creditCard, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		creditCard.Id = Id
	}
	if Number, ok := node.Properties["Number"].(string); ok {
		creditCard.Number = Number
	}
	if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
		creditCard.ExpirationDate = ExpirationDate
	}
	if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
		creditCard.SecurityCode = SecurityCode
	}
	if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
		creditCard.NameOnCard = NameOnCard
	}
	if Typ, ok := node.Properties["Typ"].(string); ok {
		creditCard.Typ = Typ
	}

	return creditCard, nil
}

func GetOnlyOneCreditCardById(id string) (CreditCard, error) {
	creditCard := CreditCard{}

	conn, err := driver.OpenPool()
	if err != nil {
		return creditCard, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ Id:{ Id } }) RETURN CreditCard", map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return creditCard, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return creditCard, NoCreditCardFound
	}
	if err != nil {
		return creditCard, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return creditCard, MultipleCreditCardFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return creditCard, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		creditCard.Id = Id
	}
	if Number, ok := node.Properties["Number"].(string); ok {
		creditCard.Number = Number
	}
	if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
		creditCard.ExpirationDate = ExpirationDate
	}
	if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
		creditCard.SecurityCode = SecurityCode
	}
	if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
		creditCard.NameOnCard = NameOnCard
	}
	if Typ, ok := node.Properties["Typ"].(string); ok {
		creditCard.Typ = Typ
	}

	return creditCard, nil
}

func GetAllCreditCardById(id string) ([]CreditCard, error) {
	var creditCards []CreditCard
	conn, err := driver.OpenPool()
	if err != nil {
		return creditCards, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ Id:{ Id } }) RETURN CreditCard", map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return creditCards, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return creditCards, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return creditCards, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		creditCard := CreditCard{}
		if Id, ok := node.Properties["Id"].(string); ok {
			creditCard.Id = Id
		}
		if Number, ok := node.Properties["Number"].(string); ok {
			creditCard.Number = Number
		}
		if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
			creditCard.ExpirationDate = ExpirationDate
		}
		if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
			creditCard.SecurityCode = SecurityCode
		}
		if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
			creditCard.NameOnCard = NameOnCard
		}
		if Typ, ok := node.Properties["Typ"].(string); ok {
			creditCard.Typ = Typ
		}

		creditCards = append(creditCards, creditCard)
	}

	return creditCards, nil
}

func UpdateAllCreditCardById(id string, creditCard CreditCard) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (creditCard:CreditCard{ Id:{ Id }) SET creditCard += { Id:{creditCardId}, Number:{creditCardNumber}, ExpirationDate:{creditCardExpirationDate}, SecurityCode:{creditCardSecurityCode}, NameOnCard:{creditCardNameOnCard}, Typ:{creditCardTyp} }", map[string]interface{}{
		"Id":                       id,
		"creditCardId":             creditCard.Id,
		"creditCardNumber":         creditCard.Number,
		"creditCardExpirationDate": creditCard.ExpirationDate,
		"creditCardSecurityCode":   creditCard.SecurityCode,
		"creditCardNameOnCard":     creditCard.NameOnCard,
		"creditCardTyp":            creditCard.Typ,
	})
	return err
}

func DeleteAllCreditCardById(id string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (creditCard:CreditCard{ Id:{ Id }) DETACH DELETE creditCard", map[string]interface{}{
		"Id": id,
	})
	return err
}

func GetCreditCardByNumber(number string) (CreditCard, error) {
	creditCard := CreditCard{}

	conn, err := driver.OpenPool()
	if err != nil {
		return creditCard, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ Number:{ Number } }) RETURN CreditCard", map[string]interface{}{
		"Number": number,
	})
	if err != nil {
		return creditCard, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return creditCard, NoCreditCardFound
	}
	if err != nil {
		return creditCard, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return creditCard, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		creditCard.Id = Id
	}
	if Number, ok := node.Properties["Number"].(string); ok {
		creditCard.Number = Number
	}
	if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
		creditCard.ExpirationDate = ExpirationDate
	}
	if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
		creditCard.SecurityCode = SecurityCode
	}
	if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
		creditCard.NameOnCard = NameOnCard
	}
	if Typ, ok := node.Properties["Typ"].(string); ok {
		creditCard.Typ = Typ
	}

	return creditCard, nil
}

func GetOnlyOneCreditCardByNumber(number string) (CreditCard, error) {
	creditCard := CreditCard{}

	conn, err := driver.OpenPool()
	if err != nil {
		return creditCard, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ Number:{ Number } }) RETURN CreditCard", map[string]interface{}{
		"Number": number,
	})

	if err != nil {
		return creditCard, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return creditCard, NoCreditCardFound
	}
	if err != nil {
		return creditCard, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return creditCard, MultipleCreditCardFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return creditCard, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		creditCard.Id = Id
	}
	if Number, ok := node.Properties["Number"].(string); ok {
		creditCard.Number = Number
	}
	if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
		creditCard.ExpirationDate = ExpirationDate
	}
	if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
		creditCard.SecurityCode = SecurityCode
	}
	if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
		creditCard.NameOnCard = NameOnCard
	}
	if Typ, ok := node.Properties["Typ"].(string); ok {
		creditCard.Typ = Typ
	}

	return creditCard, nil
}

func GetAllCreditCardByNumber(number string) ([]CreditCard, error) {
	var creditCards []CreditCard
	conn, err := driver.OpenPool()
	if err != nil {
		return creditCards, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ Number:{ Number } }) RETURN CreditCard", map[string]interface{}{
		"Number": number,
	})

	if err != nil {
		return creditCards, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return creditCards, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return creditCards, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		creditCard := CreditCard{}
		if Id, ok := node.Properties["Id"].(string); ok {
			creditCard.Id = Id
		}
		if Number, ok := node.Properties["Number"].(string); ok {
			creditCard.Number = Number
		}
		if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
			creditCard.ExpirationDate = ExpirationDate
		}
		if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
			creditCard.SecurityCode = SecurityCode
		}
		if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
			creditCard.NameOnCard = NameOnCard
		}
		if Typ, ok := node.Properties["Typ"].(string); ok {
			creditCard.Typ = Typ
		}

		creditCards = append(creditCards, creditCard)
	}

	return creditCards, nil
}

func UpdateAllCreditCardByNumber(number string, creditCard CreditCard) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (creditCard:CreditCard{ Number:{ Number }) SET creditCard += { Id:{creditCardId}, Number:{creditCardNumber}, ExpirationDate:{creditCardExpirationDate}, SecurityCode:{creditCardSecurityCode}, NameOnCard:{creditCardNameOnCard}, Typ:{creditCardTyp} }", map[string]interface{}{
		"Number":                   number,
		"creditCardId":             creditCard.Id,
		"creditCardNumber":         creditCard.Number,
		"creditCardExpirationDate": creditCard.ExpirationDate,
		"creditCardSecurityCode":   creditCard.SecurityCode,
		"creditCardNameOnCard":     creditCard.NameOnCard,
		"creditCardTyp":            creditCard.Typ,
	})
	return err
}

func DeleteAllCreditCardByNumber(number string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (creditCard:CreditCard{ Number:{ Number }) DETACH DELETE creditCard", map[string]interface{}{
		"Number": number,
	})
	return err
}

func GetCreditCardByExpirationDate(expirationDate string) (CreditCard, error) {
	creditCard := CreditCard{}

	conn, err := driver.OpenPool()
	if err != nil {
		return creditCard, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ ExpirationDate:{ ExpirationDate } }) RETURN CreditCard", map[string]interface{}{
		"ExpirationDate": expirationDate,
	})
	if err != nil {
		return creditCard, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return creditCard, NoCreditCardFound
	}
	if err != nil {
		return creditCard, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return creditCard, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		creditCard.Id = Id
	}
	if Number, ok := node.Properties["Number"].(string); ok {
		creditCard.Number = Number
	}
	if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
		creditCard.ExpirationDate = ExpirationDate
	}
	if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
		creditCard.SecurityCode = SecurityCode
	}
	if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
		creditCard.NameOnCard = NameOnCard
	}
	if Typ, ok := node.Properties["Typ"].(string); ok {
		creditCard.Typ = Typ
	}

	return creditCard, nil
}

func GetOnlyOneCreditCardByExpirationDate(expirationDate string) (CreditCard, error) {
	creditCard := CreditCard{}

	conn, err := driver.OpenPool()
	if err != nil {
		return creditCard, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ ExpirationDate:{ ExpirationDate } }) RETURN CreditCard", map[string]interface{}{
		"ExpirationDate": expirationDate,
	})

	if err != nil {
		return creditCard, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return creditCard, NoCreditCardFound
	}
	if err != nil {
		return creditCard, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return creditCard, MultipleCreditCardFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return creditCard, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		creditCard.Id = Id
	}
	if Number, ok := node.Properties["Number"].(string); ok {
		creditCard.Number = Number
	}
	if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
		creditCard.ExpirationDate = ExpirationDate
	}
	if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
		creditCard.SecurityCode = SecurityCode
	}
	if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
		creditCard.NameOnCard = NameOnCard
	}
	if Typ, ok := node.Properties["Typ"].(string); ok {
		creditCard.Typ = Typ
	}

	return creditCard, nil
}

func GetAllCreditCardByExpirationDate(expirationDate string) ([]CreditCard, error) {
	var creditCards []CreditCard
	conn, err := driver.OpenPool()
	if err != nil {
		return creditCards, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ ExpirationDate:{ ExpirationDate } }) RETURN CreditCard", map[string]interface{}{
		"ExpirationDate": expirationDate,
	})

	if err != nil {
		return creditCards, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return creditCards, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return creditCards, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		creditCard := CreditCard{}
		if Id, ok := node.Properties["Id"].(string); ok {
			creditCard.Id = Id
		}
		if Number, ok := node.Properties["Number"].(string); ok {
			creditCard.Number = Number
		}
		if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
			creditCard.ExpirationDate = ExpirationDate
		}
		if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
			creditCard.SecurityCode = SecurityCode
		}
		if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
			creditCard.NameOnCard = NameOnCard
		}
		if Typ, ok := node.Properties["Typ"].(string); ok {
			creditCard.Typ = Typ
		}

		creditCards = append(creditCards, creditCard)
	}

	return creditCards, nil
}

func UpdateAllCreditCardByExpirationDate(expirationDate string, creditCard CreditCard) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (creditCard:CreditCard{ ExpirationDate:{ ExpirationDate }) SET creditCard += { Id:{creditCardId}, Number:{creditCardNumber}, ExpirationDate:{creditCardExpirationDate}, SecurityCode:{creditCardSecurityCode}, NameOnCard:{creditCardNameOnCard}, Typ:{creditCardTyp} }", map[string]interface{}{
		"ExpirationDate":           expirationDate,
		"creditCardId":             creditCard.Id,
		"creditCardNumber":         creditCard.Number,
		"creditCardExpirationDate": creditCard.ExpirationDate,
		"creditCardSecurityCode":   creditCard.SecurityCode,
		"creditCardNameOnCard":     creditCard.NameOnCard,
		"creditCardTyp":            creditCard.Typ,
	})
	return err
}

func DeleteAllCreditCardByExpirationDate(expirationDate string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (creditCard:CreditCard{ ExpirationDate:{ ExpirationDate }) DETACH DELETE creditCard", map[string]interface{}{
		"ExpirationDate": expirationDate,
	})
	return err
}

func GetCreditCardBySecurityCode(securityCode string) (CreditCard, error) {
	creditCard := CreditCard{}

	conn, err := driver.OpenPool()
	if err != nil {
		return creditCard, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ SecurityCode:{ SecurityCode } }) RETURN CreditCard", map[string]interface{}{
		"SecurityCode": securityCode,
	})
	if err != nil {
		return creditCard, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return creditCard, NoCreditCardFound
	}
	if err != nil {
		return creditCard, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return creditCard, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		creditCard.Id = Id
	}
	if Number, ok := node.Properties["Number"].(string); ok {
		creditCard.Number = Number
	}
	if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
		creditCard.ExpirationDate = ExpirationDate
	}
	if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
		creditCard.SecurityCode = SecurityCode
	}
	if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
		creditCard.NameOnCard = NameOnCard
	}
	if Typ, ok := node.Properties["Typ"].(string); ok {
		creditCard.Typ = Typ
	}

	return creditCard, nil
}

func GetOnlyOneCreditCardBySecurityCode(securityCode string) (CreditCard, error) {
	creditCard := CreditCard{}

	conn, err := driver.OpenPool()
	if err != nil {
		return creditCard, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ SecurityCode:{ SecurityCode } }) RETURN CreditCard", map[string]interface{}{
		"SecurityCode": securityCode,
	})

	if err != nil {
		return creditCard, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return creditCard, NoCreditCardFound
	}
	if err != nil {
		return creditCard, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return creditCard, MultipleCreditCardFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return creditCard, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		creditCard.Id = Id
	}
	if Number, ok := node.Properties["Number"].(string); ok {
		creditCard.Number = Number
	}
	if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
		creditCard.ExpirationDate = ExpirationDate
	}
	if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
		creditCard.SecurityCode = SecurityCode
	}
	if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
		creditCard.NameOnCard = NameOnCard
	}
	if Typ, ok := node.Properties["Typ"].(string); ok {
		creditCard.Typ = Typ
	}

	return creditCard, nil
}

func GetAllCreditCardBySecurityCode(securityCode string) ([]CreditCard, error) {
	var creditCards []CreditCard
	conn, err := driver.OpenPool()
	if err != nil {
		return creditCards, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ SecurityCode:{ SecurityCode } }) RETURN CreditCard", map[string]interface{}{
		"SecurityCode": securityCode,
	})

	if err != nil {
		return creditCards, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return creditCards, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return creditCards, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		creditCard := CreditCard{}
		if Id, ok := node.Properties["Id"].(string); ok {
			creditCard.Id = Id
		}
		if Number, ok := node.Properties["Number"].(string); ok {
			creditCard.Number = Number
		}
		if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
			creditCard.ExpirationDate = ExpirationDate
		}
		if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
			creditCard.SecurityCode = SecurityCode
		}
		if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
			creditCard.NameOnCard = NameOnCard
		}
		if Typ, ok := node.Properties["Typ"].(string); ok {
			creditCard.Typ = Typ
		}

		creditCards = append(creditCards, creditCard)
	}

	return creditCards, nil
}

func UpdateAllCreditCardBySecurityCode(securityCode string, creditCard CreditCard) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (creditCard:CreditCard{ SecurityCode:{ SecurityCode }) SET creditCard += { Id:{creditCardId}, Number:{creditCardNumber}, ExpirationDate:{creditCardExpirationDate}, SecurityCode:{creditCardSecurityCode}, NameOnCard:{creditCardNameOnCard}, Typ:{creditCardTyp} }", map[string]interface{}{
		"SecurityCode":             securityCode,
		"creditCardId":             creditCard.Id,
		"creditCardNumber":         creditCard.Number,
		"creditCardExpirationDate": creditCard.ExpirationDate,
		"creditCardSecurityCode":   creditCard.SecurityCode,
		"creditCardNameOnCard":     creditCard.NameOnCard,
		"creditCardTyp":            creditCard.Typ,
	})
	return err
}

func DeleteAllCreditCardBySecurityCode(securityCode string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (creditCard:CreditCard{ SecurityCode:{ SecurityCode }) DETACH DELETE creditCard", map[string]interface{}{
		"SecurityCode": securityCode,
	})
	return err
}

func GetCreditCardByNameOnCard(nameOnCard string) (CreditCard, error) {
	creditCard := CreditCard{}

	conn, err := driver.OpenPool()
	if err != nil {
		return creditCard, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ NameOnCard:{ NameOnCard } }) RETURN CreditCard", map[string]interface{}{
		"NameOnCard": nameOnCard,
	})
	if err != nil {
		return creditCard, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return creditCard, NoCreditCardFound
	}
	if err != nil {
		return creditCard, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return creditCard, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		creditCard.Id = Id
	}
	if Number, ok := node.Properties["Number"].(string); ok {
		creditCard.Number = Number
	}
	if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
		creditCard.ExpirationDate = ExpirationDate
	}
	if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
		creditCard.SecurityCode = SecurityCode
	}
	if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
		creditCard.NameOnCard = NameOnCard
	}
	if Typ, ok := node.Properties["Typ"].(string); ok {
		creditCard.Typ = Typ
	}

	return creditCard, nil
}

func GetOnlyOneCreditCardByNameOnCard(nameOnCard string) (CreditCard, error) {
	creditCard := CreditCard{}

	conn, err := driver.OpenPool()
	if err != nil {
		return creditCard, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ NameOnCard:{ NameOnCard } }) RETURN CreditCard", map[string]interface{}{
		"NameOnCard": nameOnCard,
	})

	if err != nil {
		return creditCard, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return creditCard, NoCreditCardFound
	}
	if err != nil {
		return creditCard, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return creditCard, MultipleCreditCardFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return creditCard, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		creditCard.Id = Id
	}
	if Number, ok := node.Properties["Number"].(string); ok {
		creditCard.Number = Number
	}
	if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
		creditCard.ExpirationDate = ExpirationDate
	}
	if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
		creditCard.SecurityCode = SecurityCode
	}
	if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
		creditCard.NameOnCard = NameOnCard
	}
	if Typ, ok := node.Properties["Typ"].(string); ok {
		creditCard.Typ = Typ
	}

	return creditCard, nil
}

func GetAllCreditCardByNameOnCard(nameOnCard string) ([]CreditCard, error) {
	var creditCards []CreditCard
	conn, err := driver.OpenPool()
	if err != nil {
		return creditCards, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ NameOnCard:{ NameOnCard } }) RETURN CreditCard", map[string]interface{}{
		"NameOnCard": nameOnCard,
	})

	if err != nil {
		return creditCards, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return creditCards, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return creditCards, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		creditCard := CreditCard{}
		if Id, ok := node.Properties["Id"].(string); ok {
			creditCard.Id = Id
		}
		if Number, ok := node.Properties["Number"].(string); ok {
			creditCard.Number = Number
		}
		if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
			creditCard.ExpirationDate = ExpirationDate
		}
		if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
			creditCard.SecurityCode = SecurityCode
		}
		if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
			creditCard.NameOnCard = NameOnCard
		}
		if Typ, ok := node.Properties["Typ"].(string); ok {
			creditCard.Typ = Typ
		}

		creditCards = append(creditCards, creditCard)
	}

	return creditCards, nil
}

func UpdateAllCreditCardByNameOnCard(nameOnCard string, creditCard CreditCard) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (creditCard:CreditCard{ NameOnCard:{ NameOnCard }) SET creditCard += { Id:{creditCardId}, Number:{creditCardNumber}, ExpirationDate:{creditCardExpirationDate}, SecurityCode:{creditCardSecurityCode}, NameOnCard:{creditCardNameOnCard}, Typ:{creditCardTyp} }", map[string]interface{}{
		"NameOnCard":               nameOnCard,
		"creditCardId":             creditCard.Id,
		"creditCardNumber":         creditCard.Number,
		"creditCardExpirationDate": creditCard.ExpirationDate,
		"creditCardSecurityCode":   creditCard.SecurityCode,
		"creditCardNameOnCard":     creditCard.NameOnCard,
		"creditCardTyp":            creditCard.Typ,
	})
	return err
}

func DeleteAllCreditCardByNameOnCard(nameOnCard string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (creditCard:CreditCard{ NameOnCard:{ NameOnCard }) DETACH DELETE creditCard", map[string]interface{}{
		"NameOnCard": nameOnCard,
	})
	return err
}

func GetCreditCardByTyp(typ string) (CreditCard, error) {
	creditCard := CreditCard{}

	conn, err := driver.OpenPool()
	if err != nil {
		return creditCard, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ Typ:{ Typ } }) RETURN CreditCard", map[string]interface{}{
		"Typ": typ,
	})
	if err != nil {
		return creditCard, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return creditCard, NoCreditCardFound
	}
	if err != nil {
		return creditCard, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return creditCard, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		creditCard.Id = Id
	}
	if Number, ok := node.Properties["Number"].(string); ok {
		creditCard.Number = Number
	}
	if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
		creditCard.ExpirationDate = ExpirationDate
	}
	if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
		creditCard.SecurityCode = SecurityCode
	}
	if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
		creditCard.NameOnCard = NameOnCard
	}
	if Typ, ok := node.Properties["Typ"].(string); ok {
		creditCard.Typ = Typ
	}

	return creditCard, nil
}

func GetOnlyOneCreditCardByTyp(typ string) (CreditCard, error) {
	creditCard := CreditCard{}

	conn, err := driver.OpenPool()
	if err != nil {
		return creditCard, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ Typ:{ Typ } }) RETURN CreditCard", map[string]interface{}{
		"Typ": typ,
	})

	if err != nil {
		return creditCard, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return creditCard, NoCreditCardFound
	}
	if err != nil {
		return creditCard, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return creditCard, MultipleCreditCardFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return creditCard, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		creditCard.Id = Id
	}
	if Number, ok := node.Properties["Number"].(string); ok {
		creditCard.Number = Number
	}
	if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
		creditCard.ExpirationDate = ExpirationDate
	}
	if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
		creditCard.SecurityCode = SecurityCode
	}
	if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
		creditCard.NameOnCard = NameOnCard
	}
	if Typ, ok := node.Properties["Typ"].(string); ok {
		creditCard.Typ = Typ
	}

	return creditCard, nil
}

func GetAllCreditCardByTyp(typ string) ([]CreditCard, error) {
	var creditCards []CreditCard
	conn, err := driver.OpenPool()
	if err != nil {
		return creditCards, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (creditCard:CreditCard{ Typ:{ Typ } }) RETURN CreditCard", map[string]interface{}{
		"Typ": typ,
	})

	if err != nil {
		return creditCards, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return creditCards, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return creditCards, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		creditCard := CreditCard{}
		if Id, ok := node.Properties["Id"].(string); ok {
			creditCard.Id = Id
		}
		if Number, ok := node.Properties["Number"].(string); ok {
			creditCard.Number = Number
		}
		if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
			creditCard.ExpirationDate = ExpirationDate
		}
		if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
			creditCard.SecurityCode = SecurityCode
		}
		if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
			creditCard.NameOnCard = NameOnCard
		}
		if Typ, ok := node.Properties["Typ"].(string); ok {
			creditCard.Typ = Typ
		}

		creditCards = append(creditCards, creditCard)
	}

	return creditCards, nil
}

func UpdateAllCreditCardByTyp(typ string, creditCard CreditCard) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (creditCard:CreditCard{ Typ:{ Typ }) SET creditCard += { Id:{creditCardId}, Number:{creditCardNumber}, ExpirationDate:{creditCardExpirationDate}, SecurityCode:{creditCardSecurityCode}, NameOnCard:{creditCardNameOnCard}, Typ:{creditCardTyp} }", map[string]interface{}{
		"Typ":                      typ,
		"creditCardId":             creditCard.Id,
		"creditCardNumber":         creditCard.Number,
		"creditCardExpirationDate": creditCard.ExpirationDate,
		"creditCardSecurityCode":   creditCard.SecurityCode,
		"creditCardNameOnCard":     creditCard.NameOnCard,
		"creditCardTyp":            creditCard.Typ,
	})
	return err
}

func DeleteAllCreditCardByTyp(typ string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (creditCard:CreditCard{ Typ:{ Typ }) DETACH DELETE creditCard", map[string]interface{}{
		"Typ": typ,
	})
	return err
}

func GetCreditCardByCustom(query map[string]interface{}) (CreditCard, error) {
	creditCard := CreditCard{}

	conn, err := driver.OpenPool()
	if err != nil {
		return creditCard, err
	}
	defer conn.Close()

	queryStr := "MATCH (creditCard:CreditCard{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN creditCard"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return creditCard, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return creditCard, NoCreditCardFound
	}
	if err != nil {
		return creditCard, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return creditCard, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		creditCard.Id = Id
	}
	if Number, ok := node.Properties["Number"].(string); ok {
		creditCard.Number = Number
	}
	if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
		creditCard.ExpirationDate = ExpirationDate
	}
	if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
		creditCard.SecurityCode = SecurityCode
	}
	if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
		creditCard.NameOnCard = NameOnCard
	}
	if Typ, ok := node.Properties["Typ"].(string); ok {
		creditCard.Typ = Typ
	}

	return creditCard, nil
}

func GetOnlyOneCreditCardByCustom(query map[string]interface{}) (CreditCard, error) {
	creditCard := CreditCard{}

	conn, err := driver.OpenPool()
	if err != nil {
		return creditCard, err
	}
	defer conn.Close()

	queryStr := "MATCH (creditCard:CreditCard{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN creditCard"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return creditCard, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return creditCard, NoCreditCardFound
	}
	if err != nil {
		return creditCard, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return creditCard, MultipleCreditCardFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return creditCard, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		creditCard.Id = Id
	}
	if Number, ok := node.Properties["Number"].(string); ok {
		creditCard.Number = Number
	}
	if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
		creditCard.ExpirationDate = ExpirationDate
	}
	if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
		creditCard.SecurityCode = SecurityCode
	}
	if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
		creditCard.NameOnCard = NameOnCard
	}
	if Typ, ok := node.Properties["Typ"].(string); ok {
		creditCard.Typ = Typ
	}

	return creditCard, nil
}

func GetAllCreditCardByCustom(query map[string]interface{}) ([]CreditCard, error) {
	var creditCards []CreditCard

	conn, err := driver.OpenPool()
	if err != nil {
		return creditCards, err
	}
	defer conn.Close()

	queryStr := "MATCH (creditCard:CreditCard{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN creditCard"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return creditCards, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return creditCards, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return creditCards, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		creditCard := CreditCard{}
		if Id, ok := node.Properties["Id"].(string); ok {
			creditCard.Id = Id
		}
		if Number, ok := node.Properties["Number"].(string); ok {
			creditCard.Number = Number
		}
		if ExpirationDate, ok := node.Properties["ExpirationDate"].(string); ok {
			creditCard.ExpirationDate = ExpirationDate
		}
		if SecurityCode, ok := node.Properties["SecurityCode"].(string); ok {
			creditCard.SecurityCode = SecurityCode
		}
		if NameOnCard, ok := node.Properties["NameOnCard"].(string); ok {
			creditCard.NameOnCard = NameOnCard
		}
		if Typ, ok := node.Properties["Typ"].(string); ok {
			creditCard.Typ = Typ
		}

		creditCards = append(creditCards, creditCard)
	}

	return creditCards, nil
}

func UpdateAllCreditCardByCustom(params map[string]interface{}, creditCard CreditCard) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (creditCard:CreditCard{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) SET creditCard += { Id:{creditCardId}, Number:{creditCardNumber}, ExpirationDate:{creditCardExpirationDate}, SecurityCode:{creditCardSecurityCode}, NameOnCard:{creditCardNameOnCard}, Typ:{creditCardTyp} }"

	params["creditCardId"] = creditCard.Id
	params["creditCardNumber"] = creditCard.Number
	params["creditCardExpirationDate"] = creditCard.ExpirationDate
	params["creditCardSecurityCode"] = creditCard.SecurityCode
	params["creditCardNameOnCard"] = creditCard.NameOnCard
	params["creditCardTyp"] = creditCard.Typ

	_, err = conn.ExecNeo(queryStr, params)
	return err
}

func DeleteAllCreditCardByCustom(params map[string]interface{}) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (creditCard:CreditCard{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) DETACH DELETE creditCard"

	_, err = conn.ExecNeo(queryStr, params)
	return err
}
