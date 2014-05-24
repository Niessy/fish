package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// Droplet respresents a digitalocean droplet.
type Droplet struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	ImageID          int       `json:"image_id"`
	SizeID           int       `json:"size_id"`
	RegionID         int       `json:"region_id"`
	BackupsActive    bool      `json:"backups_active"`
	IPAddress        string    `json:"ip_address"`
	PrivateIPAddress string    `json:"private_ip_address"`
	Locked           bool      `json:"locked"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
}

// GetDroplets returns all users droplets, active or otherwise.
func GetDroplets() ([]Droplet, error) {
	query := fmt.Sprintf("%s?client_id=%s&api_key=%s",
		DropletsEndpoint,
		config.Conf.ClientID,
		config.Conf.APIKey)

	body, err := sendQuery(query)
	if err != nil {
		return nil, err
	}

	resp := struct {
		Status     string    `json:"status"`
		Droplets   []Droplet `json:"droplets"`
		ErrMessage string    `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Status == "ERROR" {
		return nil, errors.New(resp.ErrMessage)
	}

	return resp.Droplets, nil

}

// GetDroplet return an individual droplet based on the passed id.
func GetDroplet(id int) (Droplet, error) {
	query := fmt.Sprintf("%s/%d/?client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		config.Conf.ClientID,
		config.Conf.APIKey)

	body, err := sendQuery(query)
	if err != nil {
		return Droplet{}, err
	}

	resp := struct {
		Status     string  `json:"status"`
		Droplet    Droplet `json:"droplet"`
		ErrMessage string  `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return Droplet{}, err
	}

	if resp.Status == "ERROR" {
		return Droplet{}, errors.New(resp.ErrMessage)
	}

	return resp.Droplet, nil
}

// CreateDroplet creates a droplet based on based specs.
func CreateDroplet(name string, sizeID, imageID, regionID int, keys string) (Droplet, error) {
	query := fmt.Sprintf("%s/new?client_id=%s&api_key=%s&name=%s&size_id=%d&image_id=%d&region_id=%d&ssh_key_ids=%s",
		DropletsEndpoint,
		config.Conf.ClientID,
		config.Conf.APIKey,
		name,
		sizeID,
		imageID,
		regionID,
		keys)

	body, err := sendQuery(query)
	if err != nil {
		return Droplet{}, err
	}

	resp := struct {
		Status     string  `json:"status"`
		Droplet    Droplet `json:"droplet"`
		ErrMessage string  `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return Droplet{}, err
	}

	if resp.Status == "ERROR" {
		return Droplet{}, errors.New(resp.ErrMessage)
	}

	return resp.Droplet, nil
}

// DestroyDroplet destroys a droplet. CAUTION - this is irreversible.
// There may be more appropriate options.
func DestroyDroplet(id int) error {
	query := fmt.Sprintf("%s/%d/destroy/?client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		config.Conf.ClientID,
		config.Conf.APIKey)

	body, err := sendQuery(query)
	if err != nil {
		return err
	}

	resp := struct {
		Status     string `json:"status"`
		ErrMessage string `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Status == "ERROR" {
		return errors.New(resp.ErrMessage)
	}

	return nil
}

// ResizeDroplet droplet resizes a droplet. Sizes are based on
// the digitalocean sizes api.
func ResizeDroplet(id int, slug string) error {
	query := fmt.Sprintf("%s/%d/resize/?size_slug=%s&client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		slug,
		config.Conf.ClientID,
		config.Conf.APIKey)

	fmt.Println(query)

	body, err := sendQuery(query)
	if err != nil {
		return err
	}

	resp := struct {
		Status     string `json:"status"`
		ErrMessage string `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Status == "ERROR" {
		return errors.New(resp.ErrMessage)
	}

	return nil
}

// RebootDroplet reboots the a droplet. This is the preferred method
// to use if a server is not responding.
func RebootDroplet(id int) error {
	query := fmt.Sprintf("%s/%d/reboot/?client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		config.Conf.ClientID,
		config.Conf.APIKey)

	body, err := sendQuery(query)
	if err != nil {
		return err
	}

	resp := struct {
		Status     string `json:"status"`
		ErrMessage string `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Status == "ERROR" {
		return errors.New(resp.ErrMessage)
	}

	return nil
}

// RebootDroplet rebuilds a droplet with a default image. This can be
// useful if you want to use a different image but keep the ip address
// of the droplet.
func RebuildDroplet(imageID, id int) error {
	query := fmt.Sprintf("%s/%d/rebuild/?image_id=%d&client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		imageID,
		config.Conf.ClientID,
		config.Conf.APIKey)

	body, err := sendQuery(query)
	if err != nil {
		return err
	}

	resp := struct {
		Status     string `json:"status"`
		ErrMessage string `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Status == "ERROR" {
		return errors.New("Error rebuilding droplet with new image")
	}

	return nil
}

// StopDroplet powers off a running droplet, the droplet will remain
// in your account.
func StopDroplet(id int) error {
	query := fmt.Sprintf("%s/%d/power_off/?client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		config.Conf.ClientID,
		config.Conf.APIKey)

	body, err := sendQuery(query)
	if err != nil {
		return err
	}

	resp := struct {
		Status     string `json:"status"`
		ErrMessage string `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Status == "ERROR" {
		return errors.New("Error powering off droplet")
	}

	return nil
}

// StartDroplet powers on a powered off droplet.
func StartDroplet(id int) error {
	query := fmt.Sprintf("%s/%d/power_on/?client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		config.Conf.ClientID,
		config.Conf.APIKey)

	body, err := sendQuery(query)
	if err != nil {
		return err
	}

	resp := struct {
		Status     string `json:"status"`
		ErrMessage string `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Status == "ERROR" {
		return errors.New(resp.ErrMessage)
	}

	return nil
}

// SnapshotDroplet allows you to take a snapshot of a droplet once it is
// powered off. Be aware this may reboot the droplet.
func SnapshotDroplet(name string, id int) error {
	query := fmt.Sprintf("%s/%d/snapshot/?name=%s&client_id=%s&api_key=%s",
		DropletsEndpoint,
		id,
		name,
		config.Conf.ClientID,
		config.Conf.APIKey)

	body, err := sendQuery(query)
	if err != nil {
		return err
	}

	resp := struct {
		Status     string `json:"status"`
		ErrMessage string `json:"error_message"`
	}{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Status == "ERROR" {
		return errors.New(resp.ErrMessage)
	}

	return nil
}
