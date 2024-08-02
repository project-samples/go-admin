# go-admin

## Architecture
![User Role Service](https://cdn-images-1.medium.com/max/800/1*Gm0dypLuYaPwGM8UzrzV7w.png)

## Business Features
- User management
- Role management
## Other Features
### Authentication
- Log in by LDAP
- After logged in, get all privileges based on roles of that user
### Middleware for Authorization
#### Features
- Integrate authorization checks into HTTP middleware
- Ensure authorization is enforced on API endpoints
- Authorization: Separate the "read" and "write" permissions for 1 role, using bitwise. For example:
  - 001 (1 in decimal) is "read" permission
  - 010 (2 in decimal) is "write" permission
  - 100 (4 in decimal) is "delete" permission
  - "read" and "write" permission will be "001 | 010 = 011" (011 is 3 in decimal)
- Some other standard features
  - [config](https://github.com/core-go/core/config): load config from yaml files
  - [health check](https://github.com/core-go/core/health): to check health of SQL 
  - [logging](https://github.com/core-go/log): can use [logrus](https://github.com/sirupsen/logrus) or [zap](https://github.com/uber-go/zap) to log, support to switch between [logrus](https://github.com/sirupsen/logrus) or [zap](https://github.com/uber-go/zap)
  - log tracing by at the [middleware](https://github.com/core-go/log/tree/main/middleware) the http request and http response
### Middleware log tracing
- Log important information about incoming requests, outgoing responses, and the operations performed by the application.
### Mybatis for GO
- My batis for GOLANG.
  - Project sample is at [go-admin](https://github.com/project-samples/go-admin). Mybatis file is here [query.xml](https://github.com/project-samples/go-admin/blob/main/configs/query.xml)

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" 
        "http://mybatis.org/dtd/mybatis-3-mapper.dtd">

<mapper namespace="mappers">
  <select id="user">
    select *
    from users
    where
    <if test="username != null">
      username like #{username} and
    </if>
    <if test="displayName != null">
      displayName like #{displayName} and
    </if>
    <if test="status != null">
      status in (#{status}) and
    </if>
    <if test="q != null">
      (username like #{q} or displayName like #{q} or email like #{q}) and
    </if>
    1 = 1
    <if test="sort != null">
      order by {sort}
    </if>
    <if test="sort == null">
      order by userId
    </if>
  </select>
</mapper>
```
