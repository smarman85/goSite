#!/bin/bash
helm package gosite -d docs/
helm repo index docs --url https://smarman85.github.io/goSite/
