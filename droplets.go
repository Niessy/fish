package main

import (
	"fmt"
	"github.com/Niessy/dogo"
	"github.com/codegangsta/cli"
	"strconv"
)

func droplets(c *cli.Context) {
	droplets, err := docli.GetDroplets()
	if err != nil {
		fatalf(err.Error())
	}
	println("Droplets:")
	for _, d := range droplets {
		fmt.Printf("%s (id: %d region: %d image_id: %d ip: %q status: %s)\n",
			d.Name,
			d.ID,
			d.RegionID,
			d.ImageID,
			d.IPAddress,
			d.Status)
	}
}

func create(c *cli.Context) {
	checkArgs(c)

	d := dogo.Droplet{
		Name:          c.Args().First(),
		ImageID:       c.Int("image"),
		SizeID:        dogo.SizesMap[c.String("size")],
		RegionID:      dogo.RegionsMap[c.String("region")],
		BackupsActive: c.BoolT("backups"),
	}
	keys := c.IntSlice("keys")
	network := c.BoolT("network")

	_, err := docli.CreateDroplet(d, keys, network)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("Successfully queued %s for creation ... \n", d.Name)
}

func destroy(c *cli.Context) {
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	err = docli.DestroyDroplet(id)
	if err != nil {
		fatalf(err.Error())
	}

	fmt.Printf("Queing droplet for deletion ... \n")
}

func resize(c *cli.Context) {
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	sizeID := dogo.SizesMap[c.String("size")]

	err = docli.ResizeDroplet(id, sizeID)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("Resizing droplet\n")
}

func reboot(c *cli.Context) {
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	err = docli.RebootDroplet(id)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("rebooting droplet\n")
}

func rebuild(c *cli.Context) {
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	image := c.Int("image")

	err = docli.RebuildDroplet(id, image)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("Rebuilding droplet\n")
}

func off(c *cli.Context) {
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	err = docli.PowerOffDroplet(id)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Println("Powered Off Droplet")
}

func on(c *cli.Context) {
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	err = docli.PowerOnDroplet(id)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Println("Powered On Droplet")
}

func snapshot(c *cli.Context) {
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	name := c.String("name")

	err = docli.SnapshotDroplet(id, name)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("Created a snapshot, name = %s\n", name)
}

func restore(c *cli.Context) {
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	image := c.Int("image")

	err = docli.RestoreDroplet(id, image)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("Restored Droplet with image id = %d", image)
}

func info(c *cli.Context) {
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	droplet, err := docli.GetDroplet(id)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("Droplet:\t %s\n", droplet.Name)
	fmt.Printf("ID:\t %d\n", droplet.ID)
	fmt.Printf("Image ID:\t %d\n", droplet.ImageID)
	fmt.Printf("Size ID:\t %d\n", droplet.SizeID)
	fmt.Printf("Region ID:\t %d\n", droplet.RegionID)
	fmt.Printf("Backups Active:\t %t\n", droplet.BackupsActive)
	fmt.Printf("IP Address:\t %q\n", droplet.IPAddress)
	fmt.Printf("Private IP Address:\t %q\n", droplet.PrivateIPAddress)
	fmt.Printf("Locked:\t %t\n", droplet.Locked)
	fmt.Printf("Status:\t %s\n", droplet.Status)
	fmt.Printf("Created At:\t %v\n", droplet.CreatedAt)
}
