package org.nmf

case class Transaction(
  sender: String,
  receiver: String,
  value: Int,
  time: Int,
  category: Option[String]
)
