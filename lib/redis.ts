import { connect } from "redis";

const redisClient = await connect({
  hostname: Deno.env.get("REDIS_HOST") || "localhost",
  port: Deno.env.get("REDIS_PORT") || 6379,
  password: Deno.env.get("REDIS_PASSWORD") || "",
  name: Deno.env.get("REDIS_NAME") || "",
  username: Deno.env.get("REDIS_USERNAME") || "",
  maxRetryCount: 10,
});
console.log(await redisClient.ping());

export { redisClient };
