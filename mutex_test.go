package belajar_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 10000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

// ================================================================================================= //

//pada case ini terjadi karena ada kita membuat struct dan bakal diakses beberapa goroutine sekaligus
// RWMutex (Read Write Mutex)
// Di Go-Lang telah disediakan struct RWMutex (Read Write Mutex) untuk menangani hal ini, dimana Mutex jenis ini memiliki dua lock, lock untuk Read dan lock untuk Write

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

//jika lock dan unlock pada kedua method diatas dimatikan maka hasilnya tidak sampai 100.000
func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance", account.GetBalance())
}

// ================================================================================================= //

// Deadlock
//Hati-hati saat membuat aplikasi yang parallel atau concurrent, masalah yang sering kita hadapi adalah Deadlock
//Deadlock adalah keadaan dimana sebuah proses goroutine saling menunggu lock sehingga tidak ada satupun goroutine yang bisa jalan
//Sekarang kita coba simulasikan proses deadlock

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Eko",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Budi",
		Balance: 1000000,
	}


	go Transfer(&user1, &user2, 100000)  // user 1 sudah lock, kemudian user2 juga lock
	go Transfer(&user2, &user1, 200000)  // ternyata user2 sudah lock juga, mau lock user1, ternyata sudah kelock juga oleh goroutine

	// terjadi karena sama2 pake goroutine dan akhirnya saling tunggu2an saat lock
	// maka terjadi deadlock

	time.Sleep(10 * time.Second)

	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance) // seharusnya 90.000
}
