package transport

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AlexCorn999/bonus-system/internal/domain"
)

// ScoringSystem выполняет GET запрос в систему расчета баллов и обновляет статус и кол-во бонусов за заказ.
func (s *APIServer) ScoringSystem() {
	// получаем номер заказа из системы если его статус не PROCESSED или INVALID
	orderID, err := s.scoringsystem.GetOrderStatus()
	if err != nil {
		logError("scoringSystem", err)
		return
	}

	// создаем ссылку для запроса GET
	addr := fmt.Sprintf("%s/api/orders/%s", s.config.ScoringSystemPort, orderID)
	resp, err := http.Get(addr)
	if err != nil {
		logError("scoringSystem", err)
		return
	}
	defer resp.Body.Close()

	// анализируем статусы ответа
	if resp.StatusCode == http.StatusOK {

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			logError("scoringSystem", err)
		}

		var orderScoring domain.ScoringSystem
		if err := json.Unmarshal(data, &orderScoring); err != nil {
			logError("scoringSystem", err)
		}

		if err := s.scoringsystem.UpdateOrder(orderScoring); err != nil {
			logError("scoringSystem", err)
		}
		s.logger.Info("scoringSystem - successful bonus accrual")

	} else if resp.StatusCode == http.StatusNoContent {
		s.logger.Info("scoringSystem - order not registered")
	} else {
		s.logger.Info("scoringSystem - another error")
	}
}
