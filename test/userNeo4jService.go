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

var NoUserFound = fmt.Errorf("no user found")
var MultipleUserFound = fmt.Errorf("multiple user found")

func AddUser(user User) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE (user:User { Id:{userId}, Name:{userName}, Email:{userEmail}, Age:{userAge}, Active:{userActive}, Happy:{userHappy} })", map[string]interface{}{
		"userId":     user.Id,
		"userName":   user.Name,
		"userEmail":  user.Email,
		"userAge":    user.Age,
		"userActive": user.Active,
		"userHappy":  user.Happy,
	})

	return err
}

func GetAllUser() ([]User, error) {
	var users []User
	conn, err := driver.OpenPool()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User) RETURN user", nil)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return users, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return users, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		user := User{}
		if Id, ok := node.Properties["Id"].(string); ok {
			user.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			user.Name = Name
		}
		if Email, ok := node.Properties["Email"].(string); ok {
			user.Email = Email
		}
		if Age, ok := node.Properties["Age"].(int); ok {
			user.Age = Age
		}
		if Active, ok := node.Properties["Active"].(bool); ok {
			user.Active = Active
		}
		if Happy, ok := node.Properties["Happy"].(bool); ok {
			user.Happy = Happy
		}

		users = append(users, user)
	}

	return users, nil
}

func IndexUserById() error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE INDEX ON :User(Id)", nil)

	return err
}

func GetUserById(id string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Id:{ Id } }) RETURN User", map[string]interface{}{
		"Id": id,
	})
	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		user.Name = Name
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Age, ok := node.Properties["Age"].(int); ok {
		user.Age = Age
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Happy, ok := node.Properties["Happy"].(bool); ok {
		user.Happy = Happy
	}

	return user, nil
}

func GetOnlyOneUserById(id string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Id:{ Id } }) RETURN User", map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return user, MultipleUserFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		user.Name = Name
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Age, ok := node.Properties["Age"].(int); ok {
		user.Age = Age
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Happy, ok := node.Properties["Happy"].(bool); ok {
		user.Happy = Happy
	}

	return user, nil
}

func GetAllUserById(id string) ([]User, error) {
	var users []User
	conn, err := driver.OpenPool()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Id:{ Id } }) RETURN User", map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return users, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return users, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return users, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		user := User{}
		if Id, ok := node.Properties["Id"].(string); ok {
			user.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			user.Name = Name
		}
		if Email, ok := node.Properties["Email"].(string); ok {
			user.Email = Email
		}
		if Age, ok := node.Properties["Age"].(int); ok {
			user.Age = Age
		}
		if Active, ok := node.Properties["Active"].(bool); ok {
			user.Active = Active
		}
		if Happy, ok := node.Properties["Happy"].(bool); ok {
			user.Happy = Happy
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateAllUserById(id string, user User) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Id:{ Id }) SET user += { Id:{userId}, Name:{userName}, Email:{userEmail}, Age:{userAge}, Active:{userActive}, Happy:{userHappy} }", map[string]interface{}{
		"Id":         id,
		"userId":     user.Id,
		"userName":   user.Name,
		"userEmail":  user.Email,
		"userAge":    user.Age,
		"userActive": user.Active,
		"userHappy":  user.Happy,
	})
	return err
}

func DeleteAllUserById(id string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Id:{ Id }) DETACH DELETE user", map[string]interface{}{
		"Id": id,
	})
	return err
}

func GetUserByName(name string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Name:{ Name } }) RETURN User", map[string]interface{}{
		"Name": name,
	})
	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		user.Name = Name
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Age, ok := node.Properties["Age"].(int); ok {
		user.Age = Age
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Happy, ok := node.Properties["Happy"].(bool); ok {
		user.Happy = Happy
	}

	return user, nil
}

func GetOnlyOneUserByName(name string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Name:{ Name } }) RETURN User", map[string]interface{}{
		"Name": name,
	})

	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return user, MultipleUserFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		user.Name = Name
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Age, ok := node.Properties["Age"].(int); ok {
		user.Age = Age
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Happy, ok := node.Properties["Happy"].(bool); ok {
		user.Happy = Happy
	}

	return user, nil
}

func GetAllUserByName(name string) ([]User, error) {
	var users []User
	conn, err := driver.OpenPool()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Name:{ Name } }) RETURN User", map[string]interface{}{
		"Name": name,
	})

	if err != nil {
		return users, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return users, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return users, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		user := User{}
		if Id, ok := node.Properties["Id"].(string); ok {
			user.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			user.Name = Name
		}
		if Email, ok := node.Properties["Email"].(string); ok {
			user.Email = Email
		}
		if Age, ok := node.Properties["Age"].(int); ok {
			user.Age = Age
		}
		if Active, ok := node.Properties["Active"].(bool); ok {
			user.Active = Active
		}
		if Happy, ok := node.Properties["Happy"].(bool); ok {
			user.Happy = Happy
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateAllUserByName(name string, user User) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Name:{ Name }) SET user += { Id:{userId}, Name:{userName}, Email:{userEmail}, Age:{userAge}, Active:{userActive}, Happy:{userHappy} }", map[string]interface{}{
		"Name":       name,
		"userId":     user.Id,
		"userName":   user.Name,
		"userEmail":  user.Email,
		"userAge":    user.Age,
		"userActive": user.Active,
		"userHappy":  user.Happy,
	})
	return err
}

func DeleteAllUserByName(name string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Name:{ Name }) DETACH DELETE user", map[string]interface{}{
		"Name": name,
	})
	return err
}

func IndexUserByEmail() error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE INDEX ON :User(Email)", nil)

	return err
}

func GetUserByEmail(email string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Email:{ Email } }) RETURN User", map[string]interface{}{
		"Email": email,
	})
	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		user.Name = Name
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Age, ok := node.Properties["Age"].(int); ok {
		user.Age = Age
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Happy, ok := node.Properties["Happy"].(bool); ok {
		user.Happy = Happy
	}

	return user, nil
}

func GetOnlyOneUserByEmail(email string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Email:{ Email } }) RETURN User", map[string]interface{}{
		"Email": email,
	})

	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return user, MultipleUserFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		user.Name = Name
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Age, ok := node.Properties["Age"].(int); ok {
		user.Age = Age
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Happy, ok := node.Properties["Happy"].(bool); ok {
		user.Happy = Happy
	}

	return user, nil
}

func GetAllUserByEmail(email string) ([]User, error) {
	var users []User
	conn, err := driver.OpenPool()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Email:{ Email } }) RETURN User", map[string]interface{}{
		"Email": email,
	})

	if err != nil {
		return users, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return users, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return users, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		user := User{}
		if Id, ok := node.Properties["Id"].(string); ok {
			user.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			user.Name = Name
		}
		if Email, ok := node.Properties["Email"].(string); ok {
			user.Email = Email
		}
		if Age, ok := node.Properties["Age"].(int); ok {
			user.Age = Age
		}
		if Active, ok := node.Properties["Active"].(bool); ok {
			user.Active = Active
		}
		if Happy, ok := node.Properties["Happy"].(bool); ok {
			user.Happy = Happy
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateAllUserByEmail(email string, user User) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Email:{ Email }) SET user += { Id:{userId}, Name:{userName}, Email:{userEmail}, Age:{userAge}, Active:{userActive}, Happy:{userHappy} }", map[string]interface{}{
		"Email":      email,
		"userId":     user.Id,
		"userName":   user.Name,
		"userEmail":  user.Email,
		"userAge":    user.Age,
		"userActive": user.Active,
		"userHappy":  user.Happy,
	})
	return err
}

func DeleteAllUserByEmail(email string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Email:{ Email }) DETACH DELETE user", map[string]interface{}{
		"Email": email,
	})
	return err
}

func GetUserByAge(age int) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Age:{ Age } }) RETURN User", map[string]interface{}{
		"Age": age,
	})
	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		user.Name = Name
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Age, ok := node.Properties["Age"].(int); ok {
		user.Age = Age
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Happy, ok := node.Properties["Happy"].(bool); ok {
		user.Happy = Happy
	}

	return user, nil
}

func GetOnlyOneUserByAge(age int) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Age:{ Age } }) RETURN User", map[string]interface{}{
		"Age": age,
	})

	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return user, MultipleUserFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		user.Name = Name
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Age, ok := node.Properties["Age"].(int); ok {
		user.Age = Age
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Happy, ok := node.Properties["Happy"].(bool); ok {
		user.Happy = Happy
	}

	return user, nil
}

func GetAllUserByAge(age int) ([]User, error) {
	var users []User
	conn, err := driver.OpenPool()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Age:{ Age } }) RETURN User", map[string]interface{}{
		"Age": age,
	})

	if err != nil {
		return users, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return users, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return users, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		user := User{}
		if Id, ok := node.Properties["Id"].(string); ok {
			user.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			user.Name = Name
		}
		if Email, ok := node.Properties["Email"].(string); ok {
			user.Email = Email
		}
		if Age, ok := node.Properties["Age"].(int); ok {
			user.Age = Age
		}
		if Active, ok := node.Properties["Active"].(bool); ok {
			user.Active = Active
		}
		if Happy, ok := node.Properties["Happy"].(bool); ok {
			user.Happy = Happy
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateAllUserByAge(age int, user User) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Age:{ Age }) SET user += { Id:{userId}, Name:{userName}, Email:{userEmail}, Age:{userAge}, Active:{userActive}, Happy:{userHappy} }", map[string]interface{}{
		"Age":        age,
		"userId":     user.Id,
		"userName":   user.Name,
		"userEmail":  user.Email,
		"userAge":    user.Age,
		"userActive": user.Active,
		"userHappy":  user.Happy,
	})
	return err
}

func DeleteAllUserByAge(age int) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Age:{ Age }) DETACH DELETE user", map[string]interface{}{
		"Age": age,
	})
	return err
}

func GetUserByActive(active bool) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Active:{ Active } }) RETURN User", map[string]interface{}{
		"Active": active,
	})
	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		user.Name = Name
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Age, ok := node.Properties["Age"].(int); ok {
		user.Age = Age
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Happy, ok := node.Properties["Happy"].(bool); ok {
		user.Happy = Happy
	}

	return user, nil
}

func GetOnlyOneUserByActive(active bool) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Active:{ Active } }) RETURN User", map[string]interface{}{
		"Active": active,
	})

	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return user, MultipleUserFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		user.Name = Name
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Age, ok := node.Properties["Age"].(int); ok {
		user.Age = Age
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Happy, ok := node.Properties["Happy"].(bool); ok {
		user.Happy = Happy
	}

	return user, nil
}

func GetAllUserByActive(active bool) ([]User, error) {
	var users []User
	conn, err := driver.OpenPool()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Active:{ Active } }) RETURN User", map[string]interface{}{
		"Active": active,
	})

	if err != nil {
		return users, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return users, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return users, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		user := User{}
		if Id, ok := node.Properties["Id"].(string); ok {
			user.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			user.Name = Name
		}
		if Email, ok := node.Properties["Email"].(string); ok {
			user.Email = Email
		}
		if Age, ok := node.Properties["Age"].(int); ok {
			user.Age = Age
		}
		if Active, ok := node.Properties["Active"].(bool); ok {
			user.Active = Active
		}
		if Happy, ok := node.Properties["Happy"].(bool); ok {
			user.Happy = Happy
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateAllUserByActive(active bool, user User) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Active:{ Active }) SET user += { Id:{userId}, Name:{userName}, Email:{userEmail}, Age:{userAge}, Active:{userActive}, Happy:{userHappy} }", map[string]interface{}{
		"Active":     active,
		"userId":     user.Id,
		"userName":   user.Name,
		"userEmail":  user.Email,
		"userAge":    user.Age,
		"userActive": user.Active,
		"userHappy":  user.Happy,
	})
	return err
}

func DeleteAllUserByActive(active bool) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Active:{ Active }) DETACH DELETE user", map[string]interface{}{
		"Active": active,
	})
	return err
}

func GetUserByHappy(happy bool) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Happy:{ Happy } }) RETURN User", map[string]interface{}{
		"Happy": happy,
	})
	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		user.Name = Name
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Age, ok := node.Properties["Age"].(int); ok {
		user.Age = Age
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Happy, ok := node.Properties["Happy"].(bool); ok {
		user.Happy = Happy
	}

	return user, nil
}

func GetOnlyOneUserByHappy(happy bool) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Happy:{ Happy } }) RETURN User", map[string]interface{}{
		"Happy": happy,
	})

	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return user, MultipleUserFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		user.Name = Name
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Age, ok := node.Properties["Age"].(int); ok {
		user.Age = Age
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Happy, ok := node.Properties["Happy"].(bool); ok {
		user.Happy = Happy
	}

	return user, nil
}

func GetAllUserByHappy(happy bool) ([]User, error) {
	var users []User
	conn, err := driver.OpenPool()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Happy:{ Happy } }) RETURN User", map[string]interface{}{
		"Happy": happy,
	})

	if err != nil {
		return users, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return users, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return users, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		user := User{}
		if Id, ok := node.Properties["Id"].(string); ok {
			user.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			user.Name = Name
		}
		if Email, ok := node.Properties["Email"].(string); ok {
			user.Email = Email
		}
		if Age, ok := node.Properties["Age"].(int); ok {
			user.Age = Age
		}
		if Active, ok := node.Properties["Active"].(bool); ok {
			user.Active = Active
		}
		if Happy, ok := node.Properties["Happy"].(bool); ok {
			user.Happy = Happy
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateAllUserByHappy(happy bool, user User) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Happy:{ Happy }) SET user += { Id:{userId}, Name:{userName}, Email:{userEmail}, Age:{userAge}, Active:{userActive}, Happy:{userHappy} }", map[string]interface{}{
		"Happy":      happy,
		"userId":     user.Id,
		"userName":   user.Name,
		"userEmail":  user.Email,
		"userAge":    user.Age,
		"userActive": user.Active,
		"userHappy":  user.Happy,
	})
	return err
}

func DeleteAllUserByHappy(happy bool) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Happy:{ Happy }) DETACH DELETE user", map[string]interface{}{
		"Happy": happy,
	})
	return err
}

func GetUserByCustom(query map[string]interface{}) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	queryStr := "MATCH (user:User{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN user"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		user.Name = Name
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Age, ok := node.Properties["Age"].(int); ok {
		user.Age = Age
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Happy, ok := node.Properties["Happy"].(bool); ok {
		user.Happy = Happy
	}

	return user, nil
}

func GetOnlyOneUserByCustom(query map[string]interface{}) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	queryStr := "MATCH (user:User{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN user"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return user, MultipleUserFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		user.Name = Name
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Age, ok := node.Properties["Age"].(int); ok {
		user.Age = Age
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Happy, ok := node.Properties["Happy"].(bool); ok {
		user.Happy = Happy
	}

	return user, nil
}

func GetAllUserByCustom(query map[string]interface{}) ([]User, error) {
	var users []User

	conn, err := driver.OpenPool()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	queryStr := "MATCH (user:User{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN user"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return users, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return users, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		user := User{}
		if Id, ok := node.Properties["Id"].(string); ok {
			user.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			user.Name = Name
		}
		if Email, ok := node.Properties["Email"].(string); ok {
			user.Email = Email
		}
		if Age, ok := node.Properties["Age"].(int); ok {
			user.Age = Age
		}
		if Active, ok := node.Properties["Active"].(bool); ok {
			user.Active = Active
		}
		if Happy, ok := node.Properties["Happy"].(bool); ok {
			user.Happy = Happy
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateAllUserByCustom(params map[string]interface{}, user User) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (user:User{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) SET user += { Id:{userId}, Name:{userName}, Email:{userEmail}, Age:{userAge}, Active:{userActive}, Happy:{userHappy} }"

	params["userId"] = user.Id
	params["userName"] = user.Name
	params["userEmail"] = user.Email
	params["userAge"] = user.Age
	params["userActive"] = user.Active
	params["userHappy"] = user.Happy

	_, err = conn.ExecNeo(queryStr, params)
	return err
}

func DeleteAllUserByCustom(params map[string]interface{}) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (user:User{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) DETACH DELETE user"

	_, err = conn.ExecNeo(queryStr, params)
	return err
}
