package org.nmf

import scala.collection.mutable.Map
import cats.effect.IO
import com.twitter.finagle.Service
import com.twitter.finagle.http.{Request, Response}
import io.finch._
import io.finch.catsEffect._
import io.finch.circe._
import io.circe.generic.auto._


case class Status(status: Boolean)

case class AddTransaction(
  sender: String,
  receiver: String,
  value: Int,
  time: Int,
  category: Option[String]
) {
  def toTransaction: Transaction = Transaction(sender, receiver, value, time, category)
}

// FIXME: replace with <: TransactionsStorage
class TransactionsService[DB <: Map[Int, List[Transaction]]](db: DB) {
  def addTransaction: Endpoint[IO, Status] = {
    post("add" :: jsonBody[AddTransaction]) { request: AddTransaction =>
      val transactions = request.toTransaction :: db.getOrElse(0, List.empty)
      db.put(0, transactions)
      Ok(Status(true))
    }
  }

  def getTransactions:  Endpoint[IO, List[Transaction]] = {
    get(
      "transactions" ::
      paramOption[String]("sender") ::
      paramOption[String]("receiver") ::
      // it's probably better to use paramOption[Timestamp] instead
      paramOption[Int]("time_range_start") ::
      paramOption[Int]("time_range_end") ::
      paramOption[String]("category")
    ) {
      // FIXME: this is totally unreadable
      (sender: Option[String],
       receiver: Option[String],
       time_range_start: Option[Int],
       time_range_end: Option[Int],
       category: Option[String]) =>

      //  FIXME: hardcoded UserId
      val transactions = db.getOrElse(0, List.empty).filter(
        transaction =>
          sender.map(s => s == transaction.sender).getOrElse(true) &&
          receiver.map(s => s == transaction.receiver).getOrElse(true) &&
          transaction.time >= time_range_start.getOrElse(Int.MinValue)  &&
          transaction.time < time_range_end.getOrElse(Int.MaxValue) &&
          category.map(s => s == transaction.category).getOrElse(true)
      )

      Ok(transactions)
    }
  }
}

object TransactionsService {
  def build[DB <: Map[Int, List[Transaction]]](db: DB): Service[Request, Response] = {
    val service = new TransactionsService(db)

    Bootstrap
      .serve[Application.Json](service.addTransaction)
      .serve[Application.Json](service.getTransactions)
      .toService
  }
}