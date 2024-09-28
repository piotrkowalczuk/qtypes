// swift-tools-version:6.0
import PackageDescription

let package = Package(
    name: "qtypes",
    platforms: [
        .macOS(.v10_15), // Specify the minimum macOS version required
        .iOS(.v13),
        .watchOS(.v6),
        .tvOS(.v13),
        .visionOS(.v1)
    ],
    products: [
        .library(
            name: "QTypes",
            targets: ["QTypes"]
        )
    ],
    dependencies: [
        .package(url: "https://github.com/apple/swift-protobuf.git", from: "1.28.1"),
        .package(url: "https://github.com/grpc/grpc-swift.git", from: "1.23.1"),
        .package(url: "https://github.com/swiftlang/swift-testing.git", branch: "main"),
    ],
    targets: [
        // Główny target dla kodu źródłowego
        .target(
            name: "QTypes",
            dependencies: [
                .product(name: "SwiftProtobuf", package: "swift-protobuf"),
                .product(name: "GRPC", package: "grpc-swift"),
            ],
            path: ".",
            exclude: [
                "go.mod",
                "go.sum",
                "qtypes",
                "qtypes.go",
                "qtypes.pb.go",
                "qtypes_test.go",
                "qtypeshttp",
                "qtypes",
                "Makefile",
                "setup.py",
                "setup.cfg",
                "MANIFEST.in",
                "README.md",
                "qtypes.proto",
                "qtypes_test.pb.swift",
            ],
            sources: [
                "qtypes.pb.swift",
            ],
            publicHeadersPath: ""
        ),
       .testTarget(
            name: "QTypesTests",
            dependencies: [
                .target(name: "QTypes"),
                .product(name: "Testing", package: "swift-testing"),
            ],
            path: ".",
            exclude: [
                "go.mod",
                "go.sum",
                "qtypes",
                "qtypes.go",
                "qtypes.pb.go",
                "qtypes_test.go",
                "qtypeshttp",
                "qtypes",
                "Makefile",
                "setup.py",
                "setup.cfg",
                "MANIFEST.in",
                "README.md",
                "qtypes.proto",
                "qtypes.pb.swift"
            ],
            sources: [
                "qtypes_test.pb.swift"
            ]
        ),
    ]
)
