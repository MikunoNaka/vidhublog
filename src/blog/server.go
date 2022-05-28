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

package blog

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
  // load html templates
  r.LoadHTMLGlob("web/templates/**/*")

  // blog.vidhukant.xyz uses /blog as root
  blog := r.Group("/") 

  posts := blog.Group("/posts")

  // fetch index page
  posts.GET("/", getPosts)
  posts.POST("/filter-by-tags", filterByTagInput)

  // fetch a post
  posts.GET("/:id", getPost)
}
