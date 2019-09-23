import scala.collection.mutable
import scala.collection.immutable
import cats.effect.IO
import com.twitter.finagle.http.filter.Cors
import com.twitter.finagle.{Http, Service}
import com.twitter.finagle.http.{Request, Response}
import com.twitter.util.{Await, Try}
import io.finch._
import io.finch.catsEffect._
import io.finch.circe._
import io.circe.generic.auto._

object Category extends Enumeration {
  type Category = Value

  val Unknown: Value = Value("unknown")
  val Grocery: Value = Value("grocery")
}

case class Transaction(sender: String,
                       receiver: String,
                       value: Int,
                       time: Int,
                       category: Option[String])

object Main extends App {
  val debugState: mutable.Map[Int, List[Transaction]] = mutable.Map(0 -> List(Transaction("a", "b", 42, 1559481217, None)))
  var db: mutable.Map[Int, List[Transaction]] = debugState //mutable.Map.empty
  
  case class Status(status: Boolean)

  case class AddTransaction(sender: String,
                            receiver: String,
                            value: Int,
                            time: Int,
                            category: Option[String]) {
    def toTransaction: Transaction = Transaction(sender, receiver, value, time, category)
  }

  case class GetTransaction(sender: Option[String],
                            receiver: Option[String],
                            time_range_start: Option[Int],
                            time_range_end: Option[Int],
                            category: Option[String]
                            )

  // HTTP POST /add
  def addTransaction: Endpoint[IO, Status] = post("add" :: jsonBody[AddTransaction]) { request: AddTransaction =>
    val transactions = request.toTransaction :: db.getOrElse(0, List.empty)
    db.put(0, transactions)
    Ok(Status(true))
  }

  def getTransactions:  Endpoint[IO, List[Transaction]] = 
    get("transactions" :: paramOption[String]("sender") :: 
      paramOption[String]("receiver") ::
      paramOption[Int]("time_range_start") :: // it's probably better to use paramOption[Timestamp] instead 
      paramOption[Int]("time_range_end") :: 
      paramOption[String]("category")) { 
      (sender: Option[String], 
       receiver: Option[String],  
       time_range_start: Option[Int], 
       time_range_end: Option[Int], 
       category: Option[String]) =>
      
      val transactions = db.getOrElse(0, List.empty).filter(transaction => 
        sender.map(s => s == transaction.sender).getOrElse(true) &&
        receiver.map(s => s == transaction.receiver).getOrElse(true) &&
        transaction.time >= time_range_start.getOrElse(Int.MinValue)  &&
        transaction.time <= time_range_end.getOrElse(Int.MaxValue) && 
        category.map(s => s == transaction.category).getOrElse(true) )
      Ok(transactions)
    }

  // HTTP GET /transactions 
  def getAllTransactions: Endpoint[IO, immutable.Map[Int, List[Transaction]]] = get("transactions") {
    Ok(db.toMap)
  }

  def service: Service[Request, Response] = Bootstrap
    .serve[Application.Json](addTransaction)
    .serve[Application.Json](getTransactions)
    .toService
  
  val policy: Cors.Policy = Cors.Policy(
    allowsOrigin = _ => Some("*"),
    allowsMethods = _ => Some(Seq("GET", "POST")),
    allowsHeaders = _ => Some(Seq("Accept"))
  )
  
  val corsService: Service[Request, Response] = new Cors.HttpFilter(policy).andThen(service)

  //val db_real: PostgresProfile.backend.Database = Database.forConfig("nmf_postgres")

  Await.ready(Http.server.serve("0.0.0.0:8081", corsService))
}