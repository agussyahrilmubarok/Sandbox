package booking

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"example.com/booking/internal/logging"
	"example.com/booking/internal/tracing"
	"github.com/agussyahrilmubarok/gox/pkg/xdiscovery/xconsul"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type MemberServiceMemberResponse struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CourseServiceMessageResponse struct {
	Message string `json:"message"`
}

//go:generate mockery --name=IClient
type IClient interface {
	FindMemberByCode(ctx context.Context, memberCode string) (*MemberServiceMemberResponse, error)
	ReserveCourseByCode(ctx context.Context, courseCode string) (*CourseServiceMessageResponse, error)
	ReleaseCourseByCode(ctx context.Context, courseCode string) (*CourseServiceMessageResponse, error)
}

type client struct {
	registry *xconsul.Registry
}

func NewClient(registry *xconsul.Registry) IClient {
	return &client{
		registry: registry,
	}
}

func (c *client) FindMemberByCode(ctx context.Context, memberCode string) (*MemberServiceMemberResponse, error) {
	ctx, span := tracing.StartSpan(ctx, "cleint.FindMemberByCode")
	defer span.End()

	log := zerolog.Ctx(ctx)

	addresses, err := c.registry.ServiceAddresses(ctx, "member-app")
	if err != nil {
		log.Error().Err(err).Msg("Failed to resolve member-app from Consul")
		return nil, err
	}
	serviceAddr := addresses[0]

	url := fmt.Sprintf("http://%s/api/v1/members/find?code=%s", serviceAddr, memberCode)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Error().Err(err).Str("member_code", memberCode).Msg("Failed to create request")
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Request-ID", logging.GetRequestID(ctx))
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	httpClient := &http.Client{Timeout: 5 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Error().Err(err).Str("member_code", memberCode).Msg("Request to member-app failed")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Warn().Str("member_code", memberCode).Int("status", resp.StatusCode).Msg("Member not found or error")
		return nil, fmt.Errorf("member-app returned status: %d", resp.StatusCode)
	}

	var member MemberServiceMemberResponse
	if err := json.NewDecoder(resp.Body).Decode(&member); err != nil {
		log.Error().Err(err).Str("member_code", memberCode).Msg("Failed to decode member-app response")
		return nil, err
	}

	return &member, nil
}

func (c *client) ReserveCourseByCode(ctx context.Context, courseCode string) (*CourseServiceMessageResponse, error) {
	ctx, span := tracing.StartSpan(ctx, "cleint.ReserveCourseByCode")
	defer span.End()

	log := zerolog.Ctx(ctx)

	addresses, err := c.registry.ServiceAddresses(ctx, "course-app")
	if err != nil {
		log.Error().Err(err).Msg("Failed to resolve course-app from Consul")
		return nil, err
	}
	serviceAddr := addresses[0]

	url := fmt.Sprintf("http://%s/api/v1/courses/reserve", serviceAddr)

	payload := map[string]string{"code": courseCode}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to create request")
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Request-ID", logging.GetRequestID(ctx))
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	httpClient := &http.Client{Timeout: 5 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Error().Err(err).Str("course_code", courseCode).Msg("Request to course-app failed")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Warn().Str("course_code", courseCode).Int("status", resp.StatusCode).Msg("Failed to reserve course")
		return nil, fmt.Errorf("course-app returned status: %d", resp.StatusCode)
	}

	var result CourseServiceMessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to decode course-app response")
		return nil, err
	}

	return &result, nil
}

func (c *client) ReleaseCourseByCode(ctx context.Context, courseCode string) (*CourseServiceMessageResponse, error) {
	ctx, span := tracing.StartSpan(ctx, "client.ReleaseCourseByCode")
	defer span.End()

	log := zerolog.Ctx(ctx)

	addresses, err := c.registry.ServiceAddresses(ctx, "course-app")
	if err != nil {
		log.Error().Err(err).Msg("Failed to resolve course-app from Consul")
		return nil, err
	}
	serviceAddr := addresses[0]

	url := fmt.Sprintf("http://%s/api/v1/courses/release", serviceAddr)

	payload := map[string]string{"code": courseCode}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to create request")
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Request-ID", logging.GetRequestID(ctx))
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	httpClient := &http.Client{Timeout: 5 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Error().Err(err).Str("course_code", courseCode).Msg("Request to course-app failed")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Warn().Str("course_code", courseCode).Int("status", resp.StatusCode).Msg("Failed to release course")
		return nil, fmt.Errorf("course-app returned status: %d", resp.StatusCode)
	}

	var result CourseServiceMessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Error().Err(err).Str("course_code", courseCode).Msg("Failed to decode course-app response")
		return nil, err
	}

	return &result, nil
}
