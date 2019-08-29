import cats.effect.IO
import com.twitter.finagle.{Http, Service}
import com.twitter.finagle.http.{Request, Response}
import com.twitter.util.Await
import io.finch._
import io.finch.catsEffect._
import io.finch.circe._
import io.circe.generic.auto._

object Main extends App {

  case class Message(hello: String)

  def helloWorld: Endpoint[IO, Message] = get("hello") {
    Ok(Message("World"))
  }

  def processPost: Endpoint[IO, Message] = post("value") {
    Ok(Message("post Value"))
  }

  def service: Service[Request, Response] = Bootstrap
    .serve[Application.Json](helloWorld)
    .serve[Application.Json](processPost)
    .toService

  Await.ready(Http.server.serve(":8081", service))
}