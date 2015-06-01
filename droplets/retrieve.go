package droplets

import "github.com/digitalocean/godo"

// Get retrieve a droplet by id.
func Get(client *godo.Client, id int) (*godo.Droplet, error) {
	d, _, err := client.Droplets.Get(id)
	if err != nil {
		return nil, err
	}

	return d, nil
}