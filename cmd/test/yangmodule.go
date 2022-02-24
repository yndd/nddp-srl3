package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/openconfig/goyang/pkg/yang"
)

func parseYangModules(importDir, moduleDir []string) ([]*yang.Entry, map[string]*yang.Module, error) {
	// Read and validate the import directory with yang module

	/*
		includePaths := []string{}
		for _, path := range importDir {
			includePaths = append(includePaths, filepath.Join(path, "..."))
		}
	*/

	moduleSet := yang.NewModules()
	// Append the includePaths to the Goyang path variable, this ensures
	// that where a YANG module uses an 'include' statement to reference
	// another module, then Goyang can find this module to process.
	for _, path := range importDir {
		moduleSet.AddPath(path)
	}

	// Read the yang directory
	for _, d := range moduleDir {
		fi, err := os.Stat(d)
		if err != nil {
			return nil, nil, err
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			// Handle directory files input
			files, err := ioutil.ReadDir(d)
			if err != nil {
				return nil, nil, err
			}
			for _, f := range files {
				//g.log.Debug("Yang File Info", "FileName", d+"/"+f.Name())
				if err := moduleSet.Read(d + "/" + f.Name()); err != nil {
					return nil, nil, err
				}
			}
		case mode.IsRegular():
			// Handle file input
			//g.log.Debug("Yang File Info", "FileName", fi.Name())
			if err := moduleSet.Read(filepath.Dir(d) + fi.Name()); err != nil {
				return nil, nil, err
				//continue
			}
		}
	}

	// Process the yang modules
	errs := moduleSet.Process()
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		//return nil, nil, errs[0]
	}
	// Keep track of the top level modules we read in.
	// Those are the only modules we want to process.
	mods := map[string]*yang.Module{}
	var names []string
	for _, m := range moduleSet.Modules {
		if mods[m.Name] == nil {
			mods[m.Name] = m
			names = append(names, m.Name)
		}
	}
	sort.Strings(names)
	entries := make([]*yang.Entry, len(names))
	for x, n := range names {
		entries[x] = yang.ToEntry(mods[n])
	}
	return entries, mods, nil
}
