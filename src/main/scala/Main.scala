import scala.collection.mutable.Map
import cats.effect.IO
import com.twitter.finagle.{Http, Service}
import com.twitter.finagle.http.{Request, Response}
import com.twitter.util.{Await, Try}
import io.finch._
import io.finch.catsEffect._
import io.finch.circe._
import io.circe.generic.auto._

import scala.collection.mutable

object Category extends Enumeration {
  type Category = Value

  val Unknown: Value = Value("unknown")
  val Grocery: Value = Value("grocery")
}

case class Transaction(sender: String,
                       receiver: String,
                       value: Int,
                       time: String,
                       category: Option[String])

object Main extends App {
  var db: mutable.Map[Int, Transaction] = mutable.Map.empty

  case class Status(status: Boolean)

  case class AddTransaction(sender: String,
                            receiver: String,
                            value: Int,
                            time: String,
                            category: Option[String]) {
    def toTransaction: Transaction = Transaction(sender, receiver, value, time, category)
  }

  def addTransaction: Endpoint[IO, Status] = post("add" :: jsonBody[AddTransaction]) { request: AddTransaction =>
    db.put(0, request.toTransaction)
    Ok(Status(true))
  }

  def service: Service[Request, Response] = Bootstrap
    .serve[Application.Json](addTransaction)
    .toService

  Await.ready(Http.server.serve(":8081", service))
}