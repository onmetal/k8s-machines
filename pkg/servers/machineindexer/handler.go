/*
 * Copyright (c) 2020 by The metal-stack Authors.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package machineindexer

import (
	"fmt"
	"net/http"

	"github.com/gardener/controller-manager-library/pkg/controllermanager/server"
	"github.com/gardener/controller-manager-library/pkg/resources"
)

const CONTENT_TYPE = "Content-Type"

type requesthandler struct {
	server.Interface
	indexers []Interface
}

func (this *requesthandler) Setup() error {
	for _, t := range defaultRegistry.handlers {
		h, err := t(this)
		if err != nil {
			return err
		}
		this.indexers = append(this.indexers, h)
		err = h.Setup()
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *requesthandler) MachineIds(r *http.Request) ([]string, []string) {
	values := r.URL.Query()
	this.Infof("  found uuids: %v", values["uuid"])
	this.Infof("  found macs : %v", values["mac"])
	return values["uuid"], values["mac"]
}

func (this *requesthandler) ObjectResponse(w http.ResponseWriter, n resources.ObjectName) {
	r := fmt.Sprintf("{ \"name\": \"%s\", \"namespace\": \"%s\" }", n.Name(), n.Namespace())
	w.Write([]byte(r))
}
