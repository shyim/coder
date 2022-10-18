import { Story } from "@storybook/react"
import { MockWorkspace } from "../../testHelpers/renderHelpers"
import { AppLink, AppLinkProps } from "./AppLink"

export default {
  title: "components/AppLink",
  component: AppLink,
}

const Template: Story<AppLinkProps> = (args) => <AppLink {...args} />

export const WithIcon = Template.bind({})
WithIcon.args = {
  username: "developer",
  workspaceName: MockWorkspace.name,
  appSlug: "code-server",
  appDisplayName: "code-server",
  appIcon: "/icon/code.svg",
  appSharingLevel: "owner",
  health: "healthy",
}

export const WithoutIcon = Template.bind({})
WithoutIcon.args = {
  username: "developer",
  workspaceName: MockWorkspace.name,
  appSlug: "code-server",
  appDisplayName: "code-server",
  appSharingLevel: "owner",
  health: "healthy",
}

export const HealthDisabled = Template.bind({})
HealthDisabled.args = {
  username: "developer",
  workspaceName: MockWorkspace.name,
  appSlug: "code-server",
  appDisplayName: "code-server",
  appSharingLevel: "owner",
  health: "disabled",
}

export const HealthInitializing = Template.bind({})
HealthInitializing.args = {
  username: "developer",
  workspaceName: MockWorkspace.name,
  appSlug: "code-server",
  appDisplayName: "code-server",
  health: "initializing",
}

export const HealthUnhealthy = Template.bind({})
HealthUnhealthy.args = {
  username: "developer",
  workspaceName: MockWorkspace.name,
  appSlug: "code-server",
  appDisplayName: "code-server",
  health: "unhealthy",
}
