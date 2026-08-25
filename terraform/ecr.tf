resource "aws_ecr_repository" "portfolio" {
  name                 = "portfolio"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_iam_user" "ecr_pull" {
  name = "portfolio-ecr-pull"
}

resource "aws_iam_user_policy" "ecr_pull" {
  name = "ecr-pull-only"
  user = aws_iam_user.ecr_pull.name

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = "ecr:GetAuthorizationToken"
        Resource = "*"
      },
      {
        Effect = "Allow"
        Action = [
          "ecr:BatchCheckLayerAvailability",
          "ecr:GetDownloadUrlForLayer",
          "ecr:BatchGetImage"
        ]
        Resource = aws_ecr_repository.portfolio.arn
      }
    ]
  })
}

resource "aws_iam_access_key" "ecr_pull" {
  user = aws_iam_user.ecr_pull.name
}
