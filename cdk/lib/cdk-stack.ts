import { Stack, StackProps } from "aws-cdk-lib";
import { Construct } from "constructs";
import * as iam from "aws-cdk-lib/aws-iam";
import * as lambda from "aws-cdk-lib/aws-lambda";
import * as path from "path";

export class CdkStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    const fn = new lambda.Function(this, "ListFunc", {
      code: lambda.AssetCode.fromAsset(
        path.join(__dirname, "../../build/list.zip"),
      ),
      runtime: lambda.Runtime.GO_1_X,
      handler: "list",
    });

    fn.addToRolePolicy(
      new iam.PolicyStatement({
        effect: iam.Effect.ALLOW,
        actions: ["s3:ListBucket"],
        resources: ["*"],
      }),
    );
  }
}
