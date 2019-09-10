import slick.driver.PostgresProfile.api._
import slick.lifted.{ProvenShape, ForeignKeyQuery}
import java.sql.Timestamp

// TODO: extends to Table[Transaction]
class Transactions(tag: Tag)
  extends Table[Transaction](tag, "transactions") {

  def user_id: Rep[Int] = column[Int]("user_id")
  def sender: Rep[String] = column[String]("sender")
  def receiver: Rep[String] = column[String]("receiver")
  def value: Rep[Int] = column[Int]("value")
  def time: Rep[Timestamp] = column[Timestamp]("time")
  def category: Rep[String] = column[String]("category")

  def * : ProvenShape[Transaction] =
    (user_id, sender, receiver, value, time, category)

}

