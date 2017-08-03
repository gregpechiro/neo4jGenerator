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

var NoAddressFound = fmt.Errorf("no address found")
var MultipleAddressFound = fmt.Errorf("multiple address found")

func AddAddress(address Address) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE (address:Address { Id:{addressId}, Street:{addressStreet}, City:{addressCity}, State:{addressState}, zip:{addresszip} })", map[string]interface{}{
		"addressId":     address.Id,
		"addressStreet": address.Street,
		"addressCity":   address.City,
		"addressState":  address.State,
		"addresszip":    address.zip,
	})

	return err
}

func GetAllAddress() ([]Address, error) {
	var addresss []Address
	conn, err := driver.OpenPool()
	if err != nil {
		return addresss, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (address:Address) RETURN address", nil)
	if err != nil {
		return addresss, err
	}
	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return addresss, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return addresss, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		address := Address{}
		if Id, ok := node.Properties["Id"].(string); ok {
			address.Id = Id
		}
		if Street, ok := node.Properties["Street"].(string); ok {
			address.Street = Street
		}
		if City, ok := node.Properties["City"].(string); ok {
			address.City = City
		}
		if State, ok := node.Properties["State"].(string); ok {
			address.State = State
		}
		if zip, ok := node.Properties["zip"].(string); ok {
			address.zip = zip
		}

		addresss = append(addresss, address)
	}

	return addresss, nil
}

func GetAddressById(id string) (Address, error) {
	address := Address{}

	conn, err := driver.OpenPool()
	if err != nil {
		return address, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (address:Address{ Id:{ Id } }) RETURN Address", map[string]interface{}{
		"Id": id,
	})
	if err != nil {
		return address, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return address, NoAddressFound
	}
	if err != nil {
		return address, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return address, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		address.Id = Id
	}
	if Street, ok := node.Properties["Street"].(string); ok {
		address.Street = Street
	}
	if City, ok := node.Properties["City"].(string); ok {
		address.City = City
	}
	if State, ok := node.Properties["State"].(string); ok {
		address.State = State
	}
	if zip, ok := node.Properties["zip"].(string); ok {
		address.zip = zip
	}

	return address, nil
}

func GetOnlyOneAddressById(id string) (Address, error) {
	address := Address{}

	conn, err := driver.OpenPool()
	if err != nil {
		return address, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (address:Address{ Id:{ Id } }) RETURN Address", map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return address, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return address, NoAddressFound
	}
	if err != nil {
		return address, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return address, MultipleAddressFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return address, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		address.Id = Id
	}
	if Street, ok := node.Properties["Street"].(string); ok {
		address.Street = Street
	}
	if City, ok := node.Properties["City"].(string); ok {
		address.City = City
	}
	if State, ok := node.Properties["State"].(string); ok {
		address.State = State
	}
	if zip, ok := node.Properties["zip"].(string); ok {
		address.zip = zip
	}

	return address, nil
}

func GetAllAddressById(id string) ([]Address, error) {
	var addresss []Address
	conn, err := driver.OpenPool()
	if err != nil {
		return addresss, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (address:Address{ Id:{ Id } }) RETURN Address", map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return addresss, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return addresss, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return addresss, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		address := Address{}
		if Id, ok := node.Properties["Id"].(string); ok {
			address.Id = Id
		}
		if Street, ok := node.Properties["Street"].(string); ok {
			address.Street = Street
		}
		if City, ok := node.Properties["City"].(string); ok {
			address.City = City
		}
		if State, ok := node.Properties["State"].(string); ok {
			address.State = State
		}
		if zip, ok := node.Properties["zip"].(string); ok {
			address.zip = zip
		}

		addresss = append(addresss, address)
	}

	return addresss, nil
}

func UpdateAllAddressById(id string, address Address) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (address:Address{ Id:{ Id }) SET address += { Id:{addressId}, Street:{addressStreet}, City:{addressCity}, State:{addressState}, zip:{addresszip} }", map[string]interface{}{
		"Id":            id,
		"addressId":     address.Id,
		"addressStreet": address.Street,
		"addressCity":   address.City,
		"addressState":  address.State,
		"addresszip":    address.zip,
	})
	return err
}

func DeleteAllAddressById(id string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (address:Address{ Id:{ Id }) DETACH DELETE address", map[string]interface{}{
		"Id": id,
	})
	return err
}

func GetAddressByStreet(street string) (Address, error) {
	address := Address{}

	conn, err := driver.OpenPool()
	if err != nil {
		return address, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (address:Address{ Street:{ Street } }) RETURN Address", map[string]interface{}{
		"Street": street,
	})
	if err != nil {
		return address, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return address, NoAddressFound
	}
	if err != nil {
		return address, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return address, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		address.Id = Id
	}
	if Street, ok := node.Properties["Street"].(string); ok {
		address.Street = Street
	}
	if City, ok := node.Properties["City"].(string); ok {
		address.City = City
	}
	if State, ok := node.Properties["State"].(string); ok {
		address.State = State
	}
	if zip, ok := node.Properties["zip"].(string); ok {
		address.zip = zip
	}

	return address, nil
}

func GetOnlyOneAddressByStreet(street string) (Address, error) {
	address := Address{}

	conn, err := driver.OpenPool()
	if err != nil {
		return address, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (address:Address{ Street:{ Street } }) RETURN Address", map[string]interface{}{
		"Street": street,
	})

	if err != nil {
		return address, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return address, NoAddressFound
	}
	if err != nil {
		return address, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return address, MultipleAddressFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return address, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		address.Id = Id
	}
	if Street, ok := node.Properties["Street"].(string); ok {
		address.Street = Street
	}
	if City, ok := node.Properties["City"].(string); ok {
		address.City = City
	}
	if State, ok := node.Properties["State"].(string); ok {
		address.State = State
	}
	if zip, ok := node.Properties["zip"].(string); ok {
		address.zip = zip
	}

	return address, nil
}

func GetAllAddressByStreet(street string) ([]Address, error) {
	var addresss []Address
	conn, err := driver.OpenPool()
	if err != nil {
		return addresss, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (address:Address{ Street:{ Street } }) RETURN Address", map[string]interface{}{
		"Street": street,
	})

	if err != nil {
		return addresss, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return addresss, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return addresss, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		address := Address{}
		if Id, ok := node.Properties["Id"].(string); ok {
			address.Id = Id
		}
		if Street, ok := node.Properties["Street"].(string); ok {
			address.Street = Street
		}
		if City, ok := node.Properties["City"].(string); ok {
			address.City = City
		}
		if State, ok := node.Properties["State"].(string); ok {
			address.State = State
		}
		if zip, ok := node.Properties["zip"].(string); ok {
			address.zip = zip
		}

		addresss = append(addresss, address)
	}

	return addresss, nil
}

func UpdateAllAddressByStreet(street string, address Address) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (address:Address{ Street:{ Street }) SET address += { Id:{addressId}, Street:{addressStreet}, City:{addressCity}, State:{addressState}, zip:{addresszip} }", map[string]interface{}{
		"Street":        street,
		"addressId":     address.Id,
		"addressStreet": address.Street,
		"addressCity":   address.City,
		"addressState":  address.State,
		"addresszip":    address.zip,
	})
	return err
}

func DeleteAllAddressByStreet(street string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (address:Address{ Street:{ Street }) DETACH DELETE address", map[string]interface{}{
		"Street": street,
	})
	return err
}

func GetAddressByCity(city string) (Address, error) {
	address := Address{}

	conn, err := driver.OpenPool()
	if err != nil {
		return address, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (address:Address{ City:{ City } }) RETURN Address", map[string]interface{}{
		"City": city,
	})
	if err != nil {
		return address, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return address, NoAddressFound
	}
	if err != nil {
		return address, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return address, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		address.Id = Id
	}
	if Street, ok := node.Properties["Street"].(string); ok {
		address.Street = Street
	}
	if City, ok := node.Properties["City"].(string); ok {
		address.City = City
	}
	if State, ok := node.Properties["State"].(string); ok {
		address.State = State
	}
	if zip, ok := node.Properties["zip"].(string); ok {
		address.zip = zip
	}

	return address, nil
}

func GetOnlyOneAddressByCity(city string) (Address, error) {
	address := Address{}

	conn, err := driver.OpenPool()
	if err != nil {
		return address, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (address:Address{ City:{ City } }) RETURN Address", map[string]interface{}{
		"City": city,
	})

	if err != nil {
		return address, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return address, NoAddressFound
	}
	if err != nil {
		return address, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return address, MultipleAddressFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return address, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		address.Id = Id
	}
	if Street, ok := node.Properties["Street"].(string); ok {
		address.Street = Street
	}
	if City, ok := node.Properties["City"].(string); ok {
		address.City = City
	}
	if State, ok := node.Properties["State"].(string); ok {
		address.State = State
	}
	if zip, ok := node.Properties["zip"].(string); ok {
		address.zip = zip
	}

	return address, nil
}

func GetAllAddressByCity(city string) ([]Address, error) {
	var addresss []Address
	conn, err := driver.OpenPool()
	if err != nil {
		return addresss, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (address:Address{ City:{ City } }) RETURN Address", map[string]interface{}{
		"City": city,
	})

	if err != nil {
		return addresss, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return addresss, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return addresss, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		address := Address{}
		if Id, ok := node.Properties["Id"].(string); ok {
			address.Id = Id
		}
		if Street, ok := node.Properties["Street"].(string); ok {
			address.Street = Street
		}
		if City, ok := node.Properties["City"].(string); ok {
			address.City = City
		}
		if State, ok := node.Properties["State"].(string); ok {
			address.State = State
		}
		if zip, ok := node.Properties["zip"].(string); ok {
			address.zip = zip
		}

		addresss = append(addresss, address)
	}

	return addresss, nil
}

func UpdateAllAddressByCity(city string, address Address) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (address:Address{ City:{ City }) SET address += { Id:{addressId}, Street:{addressStreet}, City:{addressCity}, State:{addressState}, zip:{addresszip} }", map[string]interface{}{
		"City":          city,
		"addressId":     address.Id,
		"addressStreet": address.Street,
		"addressCity":   address.City,
		"addressState":  address.State,
		"addresszip":    address.zip,
	})
	return err
}

func DeleteAllAddressByCity(city string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (address:Address{ City:{ City }) DETACH DELETE address", map[string]interface{}{
		"City": city,
	})
	return err
}

func GetAddressByState(state string) (Address, error) {
	address := Address{}

	conn, err := driver.OpenPool()
	if err != nil {
		return address, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (address:Address{ State:{ State } }) RETURN Address", map[string]interface{}{
		"State": state,
	})
	if err != nil {
		return address, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return address, NoAddressFound
	}
	if err != nil {
		return address, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return address, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		address.Id = Id
	}
	if Street, ok := node.Properties["Street"].(string); ok {
		address.Street = Street
	}
	if City, ok := node.Properties["City"].(string); ok {
		address.City = City
	}
	if State, ok := node.Properties["State"].(string); ok {
		address.State = State
	}
	if zip, ok := node.Properties["zip"].(string); ok {
		address.zip = zip
	}

	return address, nil
}

func GetOnlyOneAddressByState(state string) (Address, error) {
	address := Address{}

	conn, err := driver.OpenPool()
	if err != nil {
		return address, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (address:Address{ State:{ State } }) RETURN Address", map[string]interface{}{
		"State": state,
	})

	if err != nil {
		return address, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return address, NoAddressFound
	}
	if err != nil {
		return address, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return address, MultipleAddressFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return address, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		address.Id = Id
	}
	if Street, ok := node.Properties["Street"].(string); ok {
		address.Street = Street
	}
	if City, ok := node.Properties["City"].(string); ok {
		address.City = City
	}
	if State, ok := node.Properties["State"].(string); ok {
		address.State = State
	}
	if zip, ok := node.Properties["zip"].(string); ok {
		address.zip = zip
	}

	return address, nil
}

func GetAllAddressByState(state string) ([]Address, error) {
	var addresss []Address
	conn, err := driver.OpenPool()
	if err != nil {
		return addresss, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (address:Address{ State:{ State } }) RETURN Address", map[string]interface{}{
		"State": state,
	})

	if err != nil {
		return addresss, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return addresss, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return addresss, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		address := Address{}
		if Id, ok := node.Properties["Id"].(string); ok {
			address.Id = Id
		}
		if Street, ok := node.Properties["Street"].(string); ok {
			address.Street = Street
		}
		if City, ok := node.Properties["City"].(string); ok {
			address.City = City
		}
		if State, ok := node.Properties["State"].(string); ok {
			address.State = State
		}
		if zip, ok := node.Properties["zip"].(string); ok {
			address.zip = zip
		}

		addresss = append(addresss, address)
	}

	return addresss, nil
}

func UpdateAllAddressByState(state string, address Address) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (address:Address{ State:{ State }) SET address += { Id:{addressId}, Street:{addressStreet}, City:{addressCity}, State:{addressState}, zip:{addresszip} }", map[string]interface{}{
		"State":         state,
		"addressId":     address.Id,
		"addressStreet": address.Street,
		"addressCity":   address.City,
		"addressState":  address.State,
		"addresszip":    address.zip,
	})
	return err
}

func DeleteAllAddressByState(state string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (address:Address{ State:{ State }) DETACH DELETE address", map[string]interface{}{
		"State": state,
	})
	return err
}

func GetAddressByzip(zip string) (Address, error) {
	address := Address{}

	conn, err := driver.OpenPool()
	if err != nil {
		return address, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (address:Address{ zip:{ zip } }) RETURN Address", map[string]interface{}{
		"zip": zip,
	})
	if err != nil {
		return address, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return address, NoAddressFound
	}
	if err != nil {
		return address, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return address, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		address.Id = Id
	}
	if Street, ok := node.Properties["Street"].(string); ok {
		address.Street = Street
	}
	if City, ok := node.Properties["City"].(string); ok {
		address.City = City
	}
	if State, ok := node.Properties["State"].(string); ok {
		address.State = State
	}
	if zip, ok := node.Properties["zip"].(string); ok {
		address.zip = zip
	}

	return address, nil
}

func GetOnlyOneAddressByzip(zip string) (Address, error) {
	address := Address{}

	conn, err := driver.OpenPool()
	if err != nil {
		return address, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (address:Address{ zip:{ zip } }) RETURN Address", map[string]interface{}{
		"zip": zip,
	})

	if err != nil {
		return address, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return address, NoAddressFound
	}
	if err != nil {
		return address, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return address, MultipleAddressFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return address, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		address.Id = Id
	}
	if Street, ok := node.Properties["Street"].(string); ok {
		address.Street = Street
	}
	if City, ok := node.Properties["City"].(string); ok {
		address.City = City
	}
	if State, ok := node.Properties["State"].(string); ok {
		address.State = State
	}
	if zip, ok := node.Properties["zip"].(string); ok {
		address.zip = zip
	}

	return address, nil
}

func GetAllAddressByzip(zip string) ([]Address, error) {
	var addresss []Address
	conn, err := driver.OpenPool()
	if err != nil {
		return addresss, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (address:Address{ zip:{ zip } }) RETURN Address", map[string]interface{}{
		"zip": zip,
	})

	if err != nil {
		return addresss, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return addresss, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return addresss, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		address := Address{}
		if Id, ok := node.Properties["Id"].(string); ok {
			address.Id = Id
		}
		if Street, ok := node.Properties["Street"].(string); ok {
			address.Street = Street
		}
		if City, ok := node.Properties["City"].(string); ok {
			address.City = City
		}
		if State, ok := node.Properties["State"].(string); ok {
			address.State = State
		}
		if zip, ok := node.Properties["zip"].(string); ok {
			address.zip = zip
		}

		addresss = append(addresss, address)
	}

	return addresss, nil
}

func UpdateAllAddressByzip(zip string, address Address) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (address:Address{ zip:{ zip }) SET address += { Id:{addressId}, Street:{addressStreet}, City:{addressCity}, State:{addressState}, zip:{addresszip} }", map[string]interface{}{
		"zip":           zip,
		"addressId":     address.Id,
		"addressStreet": address.Street,
		"addressCity":   address.City,
		"addressState":  address.State,
		"addresszip":    address.zip,
	})
	return err
}

func DeleteAllAddressByzip(zip string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (address:Address{ zip:{ zip }) DETACH DELETE address", map[string]interface{}{
		"zip": zip,
	})
	return err
}

func GetAddressByCustom(query map[string]interface{}) (Address, error) {
	address := Address{}

	conn, err := driver.OpenPool()
	if err != nil {
		return address, err
	}
	defer conn.Close()

	queryStr := "MATCH (address:Address{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN address"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return address, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return address, NoAddressFound
	}
	if err != nil {
		return address, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return address, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		address.Id = Id
	}
	if Street, ok := node.Properties["Street"].(string); ok {
		address.Street = Street
	}
	if City, ok := node.Properties["City"].(string); ok {
		address.City = City
	}
	if State, ok := node.Properties["State"].(string); ok {
		address.State = State
	}
	if zip, ok := node.Properties["zip"].(string); ok {
		address.zip = zip
	}

	return address, nil
}

func GetOnlyOneAddressByCustom(query map[string]interface{}) (Address, error) {
	address := Address{}

	conn, err := driver.OpenPool()
	if err != nil {
		return address, err
	}
	defer conn.Close()

	queryStr := "MATCH (address:Address{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN address"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return address, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return address, NoAddressFound
	}
	if err != nil {
		return address, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return address, MultipleAddressFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return address, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		address.Id = Id
	}
	if Street, ok := node.Properties["Street"].(string); ok {
		address.Street = Street
	}
	if City, ok := node.Properties["City"].(string); ok {
		address.City = City
	}
	if State, ok := node.Properties["State"].(string); ok {
		address.State = State
	}
	if zip, ok := node.Properties["zip"].(string); ok {
		address.zip = zip
	}

	return address, nil
}

func GetAllAddressByCustom(query map[string]interface{}) ([]Address, error) {
	var addresss []Address

	conn, err := driver.OpenPool()
	if err != nil {
		return addresss, err
	}
	defer conn.Close()

	queryStr := "MATCH (address:Address{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN address"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return addresss, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return addresss, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return addresss, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		address := Address{}
		if Id, ok := node.Properties["Id"].(string); ok {
			address.Id = Id
		}
		if Street, ok := node.Properties["Street"].(string); ok {
			address.Street = Street
		}
		if City, ok := node.Properties["City"].(string); ok {
			address.City = City
		}
		if State, ok := node.Properties["State"].(string); ok {
			address.State = State
		}
		if zip, ok := node.Properties["zip"].(string); ok {
			address.zip = zip
		}

		addresss = append(addresss, address)
	}

	return addresss, nil
}

func UpdateAllAddressByCustom(params map[string]interface{}, address Address) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (address:Address{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) SET address += { Id:{addressId}, Street:{addressStreet}, City:{addressCity}, State:{addressState}, zip:{addresszip} }"

	params["addressId"] = address.Id
	params["addressStreet"] = address.Street
	params["addressCity"] = address.City
	params["addressState"] = address.State
	params["addresszip"] = address.zip

	_, err = conn.ExecNeo(queryStr, params)
	return err
}

func DeleteAllAddressByCustom(params map[string]interface{}) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (address:Address{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) DETACH DELETE address"

	_, err = conn.ExecNeo(queryStr, params)
	return err
}
