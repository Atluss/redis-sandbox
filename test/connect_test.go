package test

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
)


func TestConn(t *testing.T) {

	ctx := context.Background()
	c, err := redis.DialContext(ctx, "tcp", ":6379")
	if err != nil {
		t.Errorf("conn to redes: %s", err)
	}
	defer c.Close()

	_, err = c.Do("SET", "foo", 1)

	if err != nil {
		t.Errorf("get arg foo: %s", err)
	}

	exists, _ := redis.Bool(c.Do("EXISTS", "foo"))
	fmt.Printf("%#v\n", exists)
}
