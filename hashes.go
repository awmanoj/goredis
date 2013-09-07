package redis

import (
	"strconv"
)

func (r *Redis) HDel(key, field string, fields ...string) (int, error) {
	args := []string{"HDEL", key, field}
	args = append(args, fields...)
	if err := r.send_command(args...); err != nil {
		return -1, err
	}
	return r.integer_reply()
}

func (r *Redis) HExists(key, field string) (bool, error) {
	if err := r.send_command("HEXISTS", key, field); err != nil {
		return false, err
	}
	return r.bool_reply()
}

func (r *Redis) HGet(key, field string) (*string, error) {
	if err := r.send_command("HGET", key, field); err != nil {
		return nil, err
	}
	return r.bulk_reply()
}

func (r *Redis) HGetAll(key string) (map[string]string, error) {
	if err := r.send_command("HGETALL", key); err != nil {
		return map[string]string{}, err
	}
	return r.stringmap_reply()
}

func (r *Redis) HIncrBy(key, field string, increment int) (int, error) {
	if err := r.send_command("HINCRBY", key, field, strconv.Itoa(increment)); err != nil {
		return -1, err
	}
	return r.integer_reply()
}

func (r *Redis) HIncrByFloat(key, field string, increment string) (string, error) {
	if err := r.send_command("HINCRBYFLOAT", key, field, increment); err != nil {
		return "", err
	}
	return r.string_reply()
}

func (r *Redis) HKeys(key string) ([]string, error) {
	if err := r.send_command("HKEYS", key); err != nil {
		return []string{}, err
	}
	return r.stringarray_reply()
}

func (r *Redis) HLen(key string) (int, error) {
	if err := r.send_command("HLEN", key); err != nil {
		return -1, err
	}
	return r.integer_reply()
}

func (r *Redis) HMGet(key, field string, fields ...string) ([]*string, error) {
	args := []string{"HMGET", key, field}
	args = append(args, fields...)
	if err := r.send_command(args...); err != nil {
		return []*string{}, err
	}
	return r.strparray_reply()
}

func (r *Redis) HMSet(key string, pairs map[string]string) error {
	if len(pairs) == 0 {
		return nil
	}
	args := []string{"HMSET", key}
	for k, v := range pairs {
		args = append(args, k, v)
	}
	if err := r.send_command(args...); err != nil {
		return err
	}
	return r.ok_reply()
}

func (r *Redis) HSet(key, field, value string) (bool, error) {
	if err := r.send_command("HSET", key, field, value); err != nil {
		return false, err
	}
	return r.bool_reply()
}

func (r *Redis) HSetnx(key, field, value string) (bool, error) {
	if err := r.send_command("HSETNX", key, field, value); err != nil {
		return false, err
	}
	return r.bool_reply()
}

func (r *Redis) HVals(key string) ([]string, error) {
	if err := r.send_command("HVALS", key); err != nil {
		return []string{}, err
	}
	return r.stringarray_reply()
}