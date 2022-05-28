/*
 * vidhublog: Vidhu Kant's Blog
 * Copyright (C) 2022  Vidhu Kant Sharma <vidhukant@protonmail.ch>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
  "github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"

  // internal routers
  "github.com/MikunoNaka/vidhublog/blog"
)

func main() {
  // initialise the router
  r := gin.New()
  r.Use(gin.Logger())
  r.Use(gin.Recovery())

  // in production nginx should handle the static files
  r.Use(static.Serve("/", static.LocalFile("./../public/", true)))

  r.GET("/x", blog.HomePage)

  blog.Routes(r)

  r.Run(":3000")
}
