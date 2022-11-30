import org.jetbrains.kotlin.gradle.tasks.KotlinCompile
import org.springframework.boot.gradle.tasks.bundling.BootBuildImage

plugins {
    id("org.springframework.boot") version "2.7.5"
    id("io.spring.dependency-management") version "1.0.15.RELEASE"
    kotlin("jvm") version "1.6.21"
    kotlin("plugin.spring") version "1.6.21"
    id("org.springframework.experimental.aot") version "0.12.1"
}

group = "com.beer.game"
version = "0.0.1-SNAPSHOT"
java.sourceCompatibility = JavaVersion.VERSION_17

repositories {
    maven { url = uri("https://repo.spring.io/release") }
    maven { url = uri("https://repo.spring.io/milestone") }
    mavenCentral()
}

val ktlint by configurations.creating

val ktlintVersion = "0.47.1"
val mockitoKotlinVersion = "2.2.0"
val springCloudVersion = "3.1.5"
dependencies {
    implementation("org.springframework.boot:spring-boot-starter-webflux")
    implementation("org.springframework.boot:spring-boot-starter-graphql")
    implementation("org.springframework.boot:spring-boot-starter-data-mongodb")
    implementation("org.springframework.boot:spring-boot-starter-data-mongodb-reactive")

    implementation("org.jetbrains.kotlin:kotlin-reflect")
    implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8")
    implementation("com.fasterxml.jackson.module:jackson-module-kotlin")

    ktlint("com.pinterest:ktlint:$ktlintVersion") {
        attributes {
            attribute(Bundling.BUNDLING_ATTRIBUTE, objects.named(Bundling.EXTERNAL))
        }
    }

    testImplementation("org.springframework.boot:spring-boot-starter-test")
    testImplementation("io.projectreactor:reactor-test")
    testImplementation("org.springframework.cloud:spring-cloud-starter-bootstrap:$springCloudVersion")
    testImplementation("org.springframework.graphql:spring-graphql-test:1.1.0")
    testImplementation("com.nhaarman.mockitokotlin2:mockito-kotlin:$mockitoKotlinVersion")
    testImplementation("com.playtika.testcontainers:embedded-mongodb:2.2.12")
}

tasks.withType<KotlinCompile> {
    kotlinOptions {
        freeCompilerArgs = listOf("-Xjsr305=strict")
        jvmTarget = "17"
    }
}

tasks.withType<Test> {
    useJUnitPlatform()
}

tasks.withType<BootBuildImage> {
    builder = "paketobuildpacks/builder:tiny"
    environment = mapOf("BP_NATIVE_IMAGE" to "true")
}


val outputDir = "${project.buildDir}/reports/ktlint/"
val inputFiles = project.fileTree(mapOf("dir" to "src", "include" to "**/*.kt"))
val ktlintCheck by tasks.creating(JavaExec::class) {
    inputs.files(inputFiles)
    outputs.dir(outputDir)

    description = "Check Kotlin code style."
    classpath = ktlint
    mainClass.set("com.pinterest.ktlint.Main")
    args = listOf("-F", "src/**/*.kt")
}

val ktlintFormat by tasks.creating(JavaExec::class) {
    inputs.files(inputFiles)
    outputs.dir(outputDir)

    description = "Fix Kotlin code style deviation."
    classpath = ktlint
    mainClass.set("com.pinterest.ktlint.Main")
    args = listOf("-F", "src/**/*.kt")
}
