Update to the latest source configuration file¶
With the upcoming deprecation of the component-config.yaml file and the endpoints.yaml file, Choreo provides a seamless process to move to the latest component.yaml version of the component.yaml file. Follow these steps to transition seamlessly:

Step 1: Identify whether your current source configuration file is outdated¶
You can determine whether you’re using an outdated configuration file in the following ways:

During a build:

On the Build page, after triggering a build, a warning will appear if your source configuration file is outdated. Click View Details in the Initialization -> Source Config Validation step. You will see an information banner with an option to update the file.
In existing deployments:

If an existing deployment uses a build with an outdated source configuration file, you will see a warning icon on the build card of the relevant environment card.
Step 2: Generate the latest component.yaml file¶
In either of the scenarios above, click Update to initiate the process of generating the latest component.yaml file. The updated file will retain your existing configurations to ensure a smooth update process.
Step 3: Add the generated configuration file to your repository¶
Download the generated component.yaml file.
Replace the existing source configuration file in the .choreo folder within your source repository.
Ensure the file name is component.yaml.
Step 4: Trigger a new build¶
Commit the new component.yaml file to your repository. Push changes to the remote Git repository.
Trigger a build using the latest commit.
Step 5: Verify the update¶
Once the build is complete:

Deploy the build.
Confirm that the warning messages no longer appear.
You can benefit from the latest features and enhancements provided by the updated source configuration file.

given by them

schemaVersion: "1.2"
endpoints:
    - name: auth-service-rest-endpoint-46b
      displayName: Auth Service REST Endpoint
      service:
        basePath: /auth-service
        port: 8080
      type: REST
      networkVisibilities:
        - Project
        - Public
      schemaFilePath: openapi.yaml
