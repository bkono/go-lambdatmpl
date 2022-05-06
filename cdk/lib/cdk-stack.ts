import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as lambda from 'aws-cdk-lib/aws-lambda'
import * as path from "path";

// import * as sqs from 'aws-cdk-lib/aws-sqs';

export class CdkStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    new lambda.Function(this, 'HelloFunc', {
      code: lambda.AssetCode.fromAsset(path.join(__dirname, "../../build/hello.zip")),
      runtime: lambda.Runtime.GO_1_X,
      handler: "hello",
    })
  }
}
