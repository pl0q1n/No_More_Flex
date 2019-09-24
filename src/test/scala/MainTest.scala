import scala.collection.mutable
import com.twitter.io.Buf
import org.scalatest.FunSuite
import io.finch._
import org.nmf.{TransactionsService, Transaction, Status}


class MainTest extends FunSuite {
  test("add transaction") {
    val db: mutable.Map[Int, List[Transaction]] = mutable.Map.empty
    val service = new TransactionsService(db)

    val request = Input.post("/add")
      .withBody[Application.Json](Buf.Utf8(
        """
        {
          "sender": "foo",
          "receiver": "bar",
          "value": 42,
          "time": 1559481217
        }
        """
      ))

    val response = service.addTransaction(request).awaitOutputUnsafe()
    assert(response.map(_.value).contains(Status(true)))
  }

  test("get transaction") {
    val db = mutable.Map(
      0 -> List(
        Transaction("a", "b", 42, 1559481217, None),
        Transaction("master", "slave", 0, 1, None)
      )
    )

    val service = new TransactionsService(db)

    {
      val request = Input.get("/transactions", "sender" -> "a")
      val response = service.getTransactions(request).awaitOutputUnsafe()
      val transactions = response.map(_.value).head
      assert(transactions.size == 1)
      assert(transactions.head.sender == "a")
    }

    {
      val request = Input.get("/transactions")
      val response = service.getTransactions(request).awaitOutputUnsafe()
      val transactions = response.map(_.value).head
      assert(transactions.size == 2)
    }
    {
      val request = Input.get("/transactions", "receiver" -> "slave")
      val response = service.getTransactions(request).awaitOutputUnsafe()
      val transactions = response.map(_.value).head
      assert(transactions.size == 1)
      assert(transactions.head.sender == "master")
    }

    {
      val request = Input.get("/transactions", "receiver" -> "romoni")
      val response = service.getTransactions(request).awaitOutputUnsafe()
      val transactions = response.map(_.value).head
      assert(transactions.size == 0)
    }
  }
}
