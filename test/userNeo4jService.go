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

	_, err = conn.ExecNeo("CREATE (user:User { id:{userid}, name:{username}, email:{useremail}, age:{userage}, active:{useractive}, happy:{userhappy} })", map[string]interface{}{
		"userid":     user.Id,
		"username":   user.Name,
		"useremail":  user.Email,
		"userage":    user.Age,
		"useractive": user.Active,
		"userhappy":  user.Happy,
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
		if id, ok := node.Properties["id"].(string); ok {
			user.Id = id
		}
		if name, ok := node.Properties["name"].(string); ok {
			user.Name = name
		}
		if email, ok := node.Properties["email"].(string); ok {
			user.Email = email
		}
		if age, ok := node.Properties["age"].(int); ok {
			user.Age = age
		}
		if active, ok := node.Properties["active"].(bool); ok {
			user.Active = active
		}
		if happy, ok := node.Properties["happy"].(bool); ok {
			user.Happy = happy
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

	_, err = conn.ExecNeo("CREATE INDEX ON :User(id)", nil)

	return err
}

func GetUserById(id string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ id:{ id } }) RETURN User", map[string]interface{}{
		"id": id,
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

	if id, ok := node.Properties["id"].(string); ok {
		user.Id = id
	}
	if name, ok := node.Properties["name"].(string); ok {
		user.Name = name
	}
	if email, ok := node.Properties["email"].(string); ok {
		user.Email = email
	}
	if age, ok := node.Properties["age"].(int); ok {
		user.Age = age
	}
	if active, ok := node.Properties["active"].(bool); ok {
		user.Active = active
	}
	if happy, ok := node.Properties["happy"].(bool); ok {
		user.Happy = happy
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

	rows, err := conn.QueryNeo("MATCH (user:User{ id:{ id } }) RETURN User", map[string]interface{}{
		"id": id,
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
	if id, ok := node.Properties["id"].(string); ok {
		user.Id = id
	}
	if name, ok := node.Properties["name"].(string); ok {
		user.Name = name
	}
	if email, ok := node.Properties["email"].(string); ok {
		user.Email = email
	}
	if age, ok := node.Properties["age"].(int); ok {
		user.Age = age
	}
	if active, ok := node.Properties["active"].(bool); ok {
		user.Active = active
	}
	if happy, ok := node.Properties["happy"].(bool); ok {
		user.Happy = happy
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

	rows, err := conn.QueryNeo("MATCH (user:User{ id:{ id } }) RETURN User", map[string]interface{}{
		"id": id,
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
		if id, ok := node.Properties["id"].(string); ok {
			user.Id = id
		}
		if name, ok := node.Properties["name"].(string); ok {
			user.Name = name
		}
		if email, ok := node.Properties["email"].(string); ok {
			user.Email = email
		}
		if age, ok := node.Properties["age"].(int); ok {
			user.Age = age
		}
		if active, ok := node.Properties["active"].(bool); ok {
			user.Active = active
		}
		if happy, ok := node.Properties["happy"].(bool); ok {
			user.Happy = happy
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

	_, err = conn.ExecNeo("MATCH (user:User{ id:{ id }) SET user += { id:{userid}, name:{username}, email:{useremail}, age:{userage}, active:{useractive}, happy:{userhappy} }", map[string]interface{}{
		"id":         id,
		"userid":     user.Id,
		"username":   user.Name,
		"useremail":  user.Email,
		"userage":    user.Age,
		"useractive": user.Active,
		"userhappy":  user.Happy,
	})
	return err
}

func DeleteAllUserById(id string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ id:{ id }) DETACH DELETE user", map[string]interface{}{
		"id": id,
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

	rows, err := conn.QueryNeo("MATCH (user:User{ name:{ name } }) RETURN User", map[string]interface{}{
		"name": name,
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

	if id, ok := node.Properties["id"].(string); ok {
		user.Id = id
	}
	if name, ok := node.Properties["name"].(string); ok {
		user.Name = name
	}
	if email, ok := node.Properties["email"].(string); ok {
		user.Email = email
	}
	if age, ok := node.Properties["age"].(int); ok {
		user.Age = age
	}
	if active, ok := node.Properties["active"].(bool); ok {
		user.Active = active
	}
	if happy, ok := node.Properties["happy"].(bool); ok {
		user.Happy = happy
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

	rows, err := conn.QueryNeo("MATCH (user:User{ name:{ name } }) RETURN User", map[string]interface{}{
		"name": name,
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
	if id, ok := node.Properties["id"].(string); ok {
		user.Id = id
	}
	if name, ok := node.Properties["name"].(string); ok {
		user.Name = name
	}
	if email, ok := node.Properties["email"].(string); ok {
		user.Email = email
	}
	if age, ok := node.Properties["age"].(int); ok {
		user.Age = age
	}
	if active, ok := node.Properties["active"].(bool); ok {
		user.Active = active
	}
	if happy, ok := node.Properties["happy"].(bool); ok {
		user.Happy = happy
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

	rows, err := conn.QueryNeo("MATCH (user:User{ name:{ name } }) RETURN User", map[string]interface{}{
		"name": name,
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
		if id, ok := node.Properties["id"].(string); ok {
			user.Id = id
		}
		if name, ok := node.Properties["name"].(string); ok {
			user.Name = name
		}
		if email, ok := node.Properties["email"].(string); ok {
			user.Email = email
		}
		if age, ok := node.Properties["age"].(int); ok {
			user.Age = age
		}
		if active, ok := node.Properties["active"].(bool); ok {
			user.Active = active
		}
		if happy, ok := node.Properties["happy"].(bool); ok {
			user.Happy = happy
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

	_, err = conn.ExecNeo("MATCH (user:User{ name:{ name }) SET user += { id:{userid}, name:{username}, email:{useremail}, age:{userage}, active:{useractive}, happy:{userhappy} }", map[string]interface{}{
		"name":       name,
		"userid":     user.Id,
		"username":   user.Name,
		"useremail":  user.Email,
		"userage":    user.Age,
		"useractive": user.Active,
		"userhappy":  user.Happy,
	})
	return err
}

func DeleteAllUserByName(name string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ name:{ name }) DETACH DELETE user", map[string]interface{}{
		"name": name,
	})
	return err
}

func IndexUserByEmail() error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE INDEX ON :User(email)", nil)

	return err
}

func GetUserByEmail(email string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ email:{ email } }) RETURN User", map[string]interface{}{
		"email": email,
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

	if id, ok := node.Properties["id"].(string); ok {
		user.Id = id
	}
	if name, ok := node.Properties["name"].(string); ok {
		user.Name = name
	}
	if email, ok := node.Properties["email"].(string); ok {
		user.Email = email
	}
	if age, ok := node.Properties["age"].(int); ok {
		user.Age = age
	}
	if active, ok := node.Properties["active"].(bool); ok {
		user.Active = active
	}
	if happy, ok := node.Properties["happy"].(bool); ok {
		user.Happy = happy
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

	rows, err := conn.QueryNeo("MATCH (user:User{ email:{ email } }) RETURN User", map[string]interface{}{
		"email": email,
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
	if id, ok := node.Properties["id"].(string); ok {
		user.Id = id
	}
	if name, ok := node.Properties["name"].(string); ok {
		user.Name = name
	}
	if email, ok := node.Properties["email"].(string); ok {
		user.Email = email
	}
	if age, ok := node.Properties["age"].(int); ok {
		user.Age = age
	}
	if active, ok := node.Properties["active"].(bool); ok {
		user.Active = active
	}
	if happy, ok := node.Properties["happy"].(bool); ok {
		user.Happy = happy
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

	rows, err := conn.QueryNeo("MATCH (user:User{ email:{ email } }) RETURN User", map[string]interface{}{
		"email": email,
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
		if id, ok := node.Properties["id"].(string); ok {
			user.Id = id
		}
		if name, ok := node.Properties["name"].(string); ok {
			user.Name = name
		}
		if email, ok := node.Properties["email"].(string); ok {
			user.Email = email
		}
		if age, ok := node.Properties["age"].(int); ok {
			user.Age = age
		}
		if active, ok := node.Properties["active"].(bool); ok {
			user.Active = active
		}
		if happy, ok := node.Properties["happy"].(bool); ok {
			user.Happy = happy
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

	_, err = conn.ExecNeo("MATCH (user:User{ email:{ email }) SET user += { id:{userid}, name:{username}, email:{useremail}, age:{userage}, active:{useractive}, happy:{userhappy} }", map[string]interface{}{
		"email":      email,
		"userid":     user.Id,
		"username":   user.Name,
		"useremail":  user.Email,
		"userage":    user.Age,
		"useractive": user.Active,
		"userhappy":  user.Happy,
	})
	return err
}

func DeleteAllUserByEmail(email string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ email:{ email }) DETACH DELETE user", map[string]interface{}{
		"email": email,
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

	rows, err := conn.QueryNeo("MATCH (user:User{ age:{ age } }) RETURN User", map[string]interface{}{
		"age": age,
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

	if id, ok := node.Properties["id"].(string); ok {
		user.Id = id
	}
	if name, ok := node.Properties["name"].(string); ok {
		user.Name = name
	}
	if email, ok := node.Properties["email"].(string); ok {
		user.Email = email
	}
	if age, ok := node.Properties["age"].(int); ok {
		user.Age = age
	}
	if active, ok := node.Properties["active"].(bool); ok {
		user.Active = active
	}
	if happy, ok := node.Properties["happy"].(bool); ok {
		user.Happy = happy
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

	rows, err := conn.QueryNeo("MATCH (user:User{ age:{ age } }) RETURN User", map[string]interface{}{
		"age": age,
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
	if id, ok := node.Properties["id"].(string); ok {
		user.Id = id
	}
	if name, ok := node.Properties["name"].(string); ok {
		user.Name = name
	}
	if email, ok := node.Properties["email"].(string); ok {
		user.Email = email
	}
	if age, ok := node.Properties["age"].(int); ok {
		user.Age = age
	}
	if active, ok := node.Properties["active"].(bool); ok {
		user.Active = active
	}
	if happy, ok := node.Properties["happy"].(bool); ok {
		user.Happy = happy
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

	rows, err := conn.QueryNeo("MATCH (user:User{ age:{ age } }) RETURN User", map[string]interface{}{
		"age": age,
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
		if id, ok := node.Properties["id"].(string); ok {
			user.Id = id
		}
		if name, ok := node.Properties["name"].(string); ok {
			user.Name = name
		}
		if email, ok := node.Properties["email"].(string); ok {
			user.Email = email
		}
		if age, ok := node.Properties["age"].(int); ok {
			user.Age = age
		}
		if active, ok := node.Properties["active"].(bool); ok {
			user.Active = active
		}
		if happy, ok := node.Properties["happy"].(bool); ok {
			user.Happy = happy
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

	_, err = conn.ExecNeo("MATCH (user:User{ age:{ age }) SET user += { id:{userid}, name:{username}, email:{useremail}, age:{userage}, active:{useractive}, happy:{userhappy} }", map[string]interface{}{
		"age":        age,
		"userid":     user.Id,
		"username":   user.Name,
		"useremail":  user.Email,
		"userage":    user.Age,
		"useractive": user.Active,
		"userhappy":  user.Happy,
	})
	return err
}

func DeleteAllUserByAge(age int) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ age:{ age }) DETACH DELETE user", map[string]interface{}{
		"age": age,
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

	rows, err := conn.QueryNeo("MATCH (user:User{ active:{ active } }) RETURN User", map[string]interface{}{
		"active": active,
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

	if id, ok := node.Properties["id"].(string); ok {
		user.Id = id
	}
	if name, ok := node.Properties["name"].(string); ok {
		user.Name = name
	}
	if email, ok := node.Properties["email"].(string); ok {
		user.Email = email
	}
	if age, ok := node.Properties["age"].(int); ok {
		user.Age = age
	}
	if active, ok := node.Properties["active"].(bool); ok {
		user.Active = active
	}
	if happy, ok := node.Properties["happy"].(bool); ok {
		user.Happy = happy
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

	rows, err := conn.QueryNeo("MATCH (user:User{ active:{ active } }) RETURN User", map[string]interface{}{
		"active": active,
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
	if id, ok := node.Properties["id"].(string); ok {
		user.Id = id
	}
	if name, ok := node.Properties["name"].(string); ok {
		user.Name = name
	}
	if email, ok := node.Properties["email"].(string); ok {
		user.Email = email
	}
	if age, ok := node.Properties["age"].(int); ok {
		user.Age = age
	}
	if active, ok := node.Properties["active"].(bool); ok {
		user.Active = active
	}
	if happy, ok := node.Properties["happy"].(bool); ok {
		user.Happy = happy
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

	rows, err := conn.QueryNeo("MATCH (user:User{ active:{ active } }) RETURN User", map[string]interface{}{
		"active": active,
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
		if id, ok := node.Properties["id"].(string); ok {
			user.Id = id
		}
		if name, ok := node.Properties["name"].(string); ok {
			user.Name = name
		}
		if email, ok := node.Properties["email"].(string); ok {
			user.Email = email
		}
		if age, ok := node.Properties["age"].(int); ok {
			user.Age = age
		}
		if active, ok := node.Properties["active"].(bool); ok {
			user.Active = active
		}
		if happy, ok := node.Properties["happy"].(bool); ok {
			user.Happy = happy
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

	_, err = conn.ExecNeo("MATCH (user:User{ active:{ active }) SET user += { id:{userid}, name:{username}, email:{useremail}, age:{userage}, active:{useractive}, happy:{userhappy} }", map[string]interface{}{
		"active":     active,
		"userid":     user.Id,
		"username":   user.Name,
		"useremail":  user.Email,
		"userage":    user.Age,
		"useractive": user.Active,
		"userhappy":  user.Happy,
	})
	return err
}

func DeleteAllUserByActive(active bool) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ active:{ active }) DETACH DELETE user", map[string]interface{}{
		"active": active,
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

	rows, err := conn.QueryNeo("MATCH (user:User{ happy:{ happy } }) RETURN User", map[string]interface{}{
		"happy": happy,
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

	if id, ok := node.Properties["id"].(string); ok {
		user.Id = id
	}
	if name, ok := node.Properties["name"].(string); ok {
		user.Name = name
	}
	if email, ok := node.Properties["email"].(string); ok {
		user.Email = email
	}
	if age, ok := node.Properties["age"].(int); ok {
		user.Age = age
	}
	if active, ok := node.Properties["active"].(bool); ok {
		user.Active = active
	}
	if happy, ok := node.Properties["happy"].(bool); ok {
		user.Happy = happy
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

	rows, err := conn.QueryNeo("MATCH (user:User{ happy:{ happy } }) RETURN User", map[string]interface{}{
		"happy": happy,
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
	if id, ok := node.Properties["id"].(string); ok {
		user.Id = id
	}
	if name, ok := node.Properties["name"].(string); ok {
		user.Name = name
	}
	if email, ok := node.Properties["email"].(string); ok {
		user.Email = email
	}
	if age, ok := node.Properties["age"].(int); ok {
		user.Age = age
	}
	if active, ok := node.Properties["active"].(bool); ok {
		user.Active = active
	}
	if happy, ok := node.Properties["happy"].(bool); ok {
		user.Happy = happy
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

	rows, err := conn.QueryNeo("MATCH (user:User{ happy:{ happy } }) RETURN User", map[string]interface{}{
		"happy": happy,
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
		if id, ok := node.Properties["id"].(string); ok {
			user.Id = id
		}
		if name, ok := node.Properties["name"].(string); ok {
			user.Name = name
		}
		if email, ok := node.Properties["email"].(string); ok {
			user.Email = email
		}
		if age, ok := node.Properties["age"].(int); ok {
			user.Age = age
		}
		if active, ok := node.Properties["active"].(bool); ok {
			user.Active = active
		}
		if happy, ok := node.Properties["happy"].(bool); ok {
			user.Happy = happy
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

	_, err = conn.ExecNeo("MATCH (user:User{ happy:{ happy }) SET user += { id:{userid}, name:{username}, email:{useremail}, age:{userage}, active:{useractive}, happy:{userhappy} }", map[string]interface{}{
		"happy":      happy,
		"userid":     user.Id,
		"username":   user.Name,
		"useremail":  user.Email,
		"userage":    user.Age,
		"useractive": user.Active,
		"userhappy":  user.Happy,
	})
	return err
}

func DeleteAllUserByHappy(happy bool) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ happy:{ happy }) DETACH DELETE user", map[string]interface{}{
		"happy": happy,
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

	if id, ok := node.Properties["id"].(string); ok {
		user.Id = id
	}
	if name, ok := node.Properties["name"].(string); ok {
		user.Name = name
	}
	if email, ok := node.Properties["email"].(string); ok {
		user.Email = email
	}
	if age, ok := node.Properties["age"].(int); ok {
		user.Age = age
	}
	if active, ok := node.Properties["active"].(bool); ok {
		user.Active = active
	}
	if happy, ok := node.Properties["happy"].(bool); ok {
		user.Happy = happy
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

	if id, ok := node.Properties["id"].(string); ok {
		user.Id = id
	}
	if name, ok := node.Properties["name"].(string); ok {
		user.Name = name
	}
	if email, ok := node.Properties["email"].(string); ok {
		user.Email = email
	}
	if age, ok := node.Properties["age"].(int); ok {
		user.Age = age
	}
	if active, ok := node.Properties["active"].(bool); ok {
		user.Active = active
	}
	if happy, ok := node.Properties["happy"].(bool); ok {
		user.Happy = happy
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
		if id, ok := node.Properties["id"].(string); ok {
			user.Id = id
		}
		if name, ok := node.Properties["name"].(string); ok {
			user.Name = name
		}
		if email, ok := node.Properties["email"].(string); ok {
			user.Email = email
		}
		if age, ok := node.Properties["age"].(int); ok {
			user.Age = age
		}
		if active, ok := node.Properties["active"].(bool); ok {
			user.Active = active
		}
		if happy, ok := node.Properties["happy"].(bool); ok {
			user.Happy = happy
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
	queryStr += "}) SET user += { id:{userid}, name:{username}, email:{useremail}, age:{userage}, active:{useractive}, happy:{userhappy} }"

	params["userid"] = user.Id
	params["username"] = user.Name
	params["useremail"] = user.Email
	params["userage"] = user.Age
	params["useractive"] = user.Active
	params["userhappy"] = user.Happy

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
