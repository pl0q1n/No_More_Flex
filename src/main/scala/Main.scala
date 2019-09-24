package org.nmf

import scala.collection.mutable
import com.twitter.finagle.http.filter.Cors
import com.twitter.finagle.{Service, Http}
import com.twitter.finagle.http.{Request, Response}
import com.twitter.finagle.transport.Transport
import com.twitter.util.Await
import com.twitter.logging.Logger


object Main extends App {
  val log = Logger.get(getClass)
  val db: mutable.Map[Int, List[Transaction]] = {
    mutable.Map(
      0 -> List(Transaction("a", "b", 42, 1559481217, None)),
      1 -> List(Transaction("b", "c", 1337, 13247980, None))
    )
  }

  val policy: Cors.Policy = Cors.Policy(
    allowsOrigin = _ => Some("*"),
    allowsMethods = _ => Some(Seq("GET", "POST")),
    allowsHeaders = _ => Some(Seq("Accept"))
  )

  val service: Service[Request, Response] = new Cors.HttpFilter(policy)
    .andThen(TransactionsService.build(db))

  val endpoint = "0.0.0.0:80"
  val server = Http.server.configured(Transport.Options(noDelay = false, reuseAddr = true))

  log.debug(s"Starting on $endpoint")
  Await.ready(server.serve(endpoint, service))
}