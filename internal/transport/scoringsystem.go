package transport

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/AlexCorn999/bonus-system/internal/domain"
)

// функция должна делать гет запросы по адресу и читать тело затем обновлять данные в базе
func (s *APIServer) ScoringSystem() {

	var usr domain.SighUpAndInInput
	usr.Login = "Alex"
	usr.Password = "12345678"

	if err := s.users.SignUp(usr); err != nil {
		fmt.Println(err)
	}

	var userID int64 = 1
	ctx := context.WithValue(context.Background(), domain.UserIDKeyForContext, userID)

	if err := s.orders.AddOrderID(ctx, "5555555555554444"); err != nil {
		fmt.Println(err)
	}
	if err := s.orders.AddOrderID(ctx, "20412011"); err != nil {
		fmt.Println(err)
	}
	user, _ := ctx.Value(domain.UserIDKeyForContext).(int64)
	fmt.Println(user)

	time.Sleep(time.Second * 10)
	addr := fmt.Sprintf("%s/api/orders/5555555555554444", s.config.ScoringSystemPort)
	resp, err := http.Get(addr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response status code:", resp.StatusCode)
	fmt.Println("Response body:", string(body))

	addr2 := fmt.Sprintf("%s/api/orders/20412011", s.config.ScoringSystemPort)
	resp2, err := http.Get(addr2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp2.Body.Close()

	body2, err := io.ReadAll(resp2.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response status code:", resp2.StatusCode)
	fmt.Println("Response body:", string(body2))
	// берем заказ из системы если его нету то ошибка 204 заказа нет в системе

	// если заказ есть то делаем гет запрос

	// читаем тело ответа и заносим в поля заказа

	// 429 — превышено количество запросов к сервису.
}
