package main

import "testing"

// TestHello 测试 Hello 函数在不同情况下的输出
func TestHello(t *testing.T) {
 t.Run("saying hello to people", func(t *testing.T) {
  got := Hello("chris", "")
  want := "Hello, chris"
  assertCorrectMessage(t, got, want)
 })

 t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
  got := Hello("", "")
  want := "Hello, world"
  assertCorrectMessage(t, got, want)
 })

 t.Run("in spanish", func(t *testing.T) {
  got := Hello("Elodie", "spanish")
  want := "Hola, Elodie"
  assertCorrectMessage(t, got, want)
 })

 t.Run("in French", func(t *testing.T) {
  got := Hello("Alice", "French")
  want := "Bonjour, Alice"
  assertCorrectMessage(t, got, want)
 })
