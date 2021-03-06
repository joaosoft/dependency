package main

import (
	"fmt"

	"os"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

type Dependency struct {
	config        *DependencyConfig
	quit          chan int
	isLogExternal bool
	pm            *manager.Manager
	logger        logger.ILogger
	vcs           *Vcs
	vendor        string
	bkVendor      string
}

func NewDependency(options ...DependencyOption) (*Dependency, error) {
	config, simpleConfig, err := NewConfig()
	log := logger.NewLogDefault("dependency", logger.WarnLevel)

	vcs, err := NewVcs(fmt.Sprintf("%s/%s", os.Getenv("HOME"), CacheRepository), CacheRepositoryConfigFile, ProtocolHTTPS, log)
	if err != nil {
		return nil, err
	}

	service := &Dependency{
		quit:   make(chan int),
		pm:     manager.NewManager(manager.WithRunInBackground(true), manager.WithLogLevel(logger.WarnLevel)),
		logger: log,
		vcs:    vcs,
		vendor: "vendor",
		config: config.Dependency,
	}

	if service.isLogExternal {
		service.pm.Reconfigure(manager.WithLogger(service.logger))
	}

	if err != nil {
		log.Error(err.Error())
	} else if config.Dependency != nil {
		service.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Dependency.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
	}

	service.Reconfigure(options...)

	return service, nil
}

func (d *Dependency) Get() error {
	d.logger.Debug("executing Get")

	var err error
	loadedImports := make(map[string]bool)
	installedImports := make(Imports)

	defer func() {
		if err != nil {
			d.doUndoBackupVendor()
		}
	}()

	// backup old vendor folder
	if err = d.doBackupVendor(); err != nil {
		return err
	}

	dir, _ := os.Getwd()
	if err = d.doGet(dir, loadedImports, installedImports, false, false); err != nil {
		return err
	} else {
		// save generated imports
		if err = d.doSaveImports(installedImports); err != nil {
			return err
		}
	}

	return nil
}

func (d *Dependency) Update() error {
	d.logger.Debug("executing Update")

	var err error
	loadedImports := make(map[string]bool)
	installedImports := make(Imports)

	defer func() {
		if err != nil {
			d.doUndoBackupVendor()
		}
	}()

	// backup old vendor folder
	if err = d.doBackupVendor(); err != nil {
		return err
	}

	if err := d.doClearGen(); err != nil {
		return err
	}

	dir, _ := os.Getwd()
	if err = d.doGet(dir, loadedImports, installedImports, false, true); err != nil {
		return err
	} else {
		// save generated imports
		if err = d.doSaveImports(installedImports); err != nil {
			return err
		}
	}

	return nil
}

func (d *Dependency) Reset() error {
	d.logger.Debug("executing Reset")

	if err := d.doClearGen(); err != nil {
		return err
	}

	if err := d.doClearLock(); err != nil {
		return err
	}

	if err := d.Get(); err != nil {
		return err
	}

	return nil
}

func (d *Dependency) Add(newImport string) error {
	d.logger.Debug("executing Add")

	var err error
	loadedImports := make(map[string]bool)
	installedImports := make(Imports)

	if err = d.doAdd(loadedImports, installedImports, newImport); err != nil {
		return err
	} else {
		// save generated imports
		if err = d.doSaveImports(installedImports); err != nil {
			return err
		}
	}

	return nil
}

func (d *Dependency) Remove(removeImport string) error {
	d.logger.Debug("executing Add")

	var err error
	loadedImports := make(map[string]bool)
	installedImports := make(Imports)

	if err = d.doRemove(loadedImports, installedImports, removeImport); err != nil {
		return err
	} else {
		// save generated imports
		if err = d.doSaveImports(installedImports); err != nil {
			return err
		}
	}

	return nil
}
