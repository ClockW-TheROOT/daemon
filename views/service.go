/* cSploit - a simple penetration testing suite
 * Copyright (C) 2016 Massimo Dragano aka tux_mind <tux_mind@csploit.org>
 *
 * cSploit is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * cSploit is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with cSploit.  If not, see <http://www.gnu.org/licenses/\>.
 *
 */
package views

import "github.com/cSploit/daemon/models"

func ServiceIndex(arg interface{}) interface{} {
	svc := arg.([]models.Service)

	return svc
}

func ServiceShow(arg interface{}) interface{} {
	svc := arg.(models.Service)

	return svc
}

func serviceAsChild(arg interface{}) interface{} {
	svc := arg.(*models.Service)

	if svc == nil {
		return ""
	} else {
		return svc.FormatName()
	}
}
