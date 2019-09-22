val finchVersion = "0.26.0"
val circeVersion = "0.10.1"
val scalatestVersion = "3.0.5"

cancelable in Global := true

enablePlugins(JavaAppPackaging)
enablePlugins(DockerPlugin)

lazy val root = (project in file("."))
  .settings(
    organization := "com.example",
    name := "no_more_flex",
    version := "0.0.1-SNAPSHOT",
    scalaVersion := "2.12.7",
    libraryDependencies ++= Seq(
      "com.github.finagle" %% "finchx-core"  % finchVersion,
      "com.github.finagle" %% "finchx-circe"  % finchVersion,
      "io.circe" %% "circe-generic" % circeVersion,
      "org.scalatest"      %% "scalatest"    % scalatestVersion % "test",
      "com.typesafe.slick" %% "slick" % "3.3.2",
      "org.postgresql" % "postgresql" % "9.3-1100-jdbc41"
    )
  )

mainClass in Compile := Some("Main")