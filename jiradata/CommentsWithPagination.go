package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// /bin/slipscheme -pkg jiradata ActorInput.json AddField.json AddGroup.json ApplicationProperty.json ApplicationRole.json Attachment.json AttachmentArchiveImpl.json AttachmentMeta.json AuditRecord.json AuthParams.json AuthSuccess.json AutoCompleteResponse.json AutoCompleteResultWrapper.json Avatar.json AvatarCropping.json BulkOperationErrorResult.json Comment.json CommentsWithPagination.json Component.json ComponentIssueCounts.json Configuration.json CreateMeta.json CreateUpdateRoleRequest.json CurrentUser.json CustomFieldDefinition.json CustomFieldOption.json Dashboard.json Dashboards.json Default.json DefaultShareScope.json DeleteAndReplaceVersion.json EditMeta.json EntityPropertiesKeys.json EntityProperty.json ErrorCollection.json Field.json Filter.json FilterPermission.json Group.json GroupSuggestions.json HumanReadableArchive.json Id.json IndexSummary.json Issue.json IssueCreateResponse.json IssueLink.json IssueLinkType.json IssueLinkTypes.json IssuePickerResult.json IssueSubTaskMovePosition.json IssueType.json IssueTypeCreate.json IssueTypeMapping.json IssueTypeUpdate.json IssueUpdate.json IssuesUpdate.json LinkIssueRequest.json ListofApplicationRole.json ListofAttachment.json ListofColumnItem.json ListofColumnLayoutItem.json ListofComponent.json ListofField.json ListofFilter.json ListofFilterPermission.json ListofIssueType.json ListofIssueTypeWithStatus.json ListofPriority.json ListofProject.json ListofProjectCategory.json ListofProjectType.json ListofProperty.json ListofRemoteEntityLink.json ListofRemoteIssueLink.json ListofResolution.json ListofScreenableTab.json ListofStatus.json ListofStatusCategory.json ListofUser.json ListofVersion.json ListofWorkflow.json ListofWorkflowMapping.json ListofWorklog.json MoveField.json Notification.json NotificationScheme.json PageofCustomField.json PageofNotificationScheme.json PageofVersion.json Password.json PasswordPolicyCreateUser.json PasswordPolicyUpdateUser.json PermissionGrant.json PermissionGrants.json PermissionScheme.json PermissionSchemeAttribute.json PermissionSchemes.json Permissions.json Priority.json PriorityScheme.json PrioritySchemeList.json PrioritySchemeUpdate.json Project.json ProjectCategory.json ProjectIdentity.json ProjectInput.json ProjectRole.json ProjectRoleActorsUpdate.json ProjectType.json Property.json Reindex.json ReindexRequest.json RemoteEntityLink.json RemoteIssueLink.json RemoteIssueLinkCreateOrUpdateRequest.json RemoteIssueLinkCreateOrUpdateResponse.json Resolution.json Response.json ScreenableField.json ScreenableTab.json SearchRequest.json SearchResults.json SecurityLevel.json SecurityListLevel.json SecurityScheme.json SecuritySchemes.json ServerInfo.json SharePermissionInput.json Status.json StatusCategory.json String.json SystemAvatars.json TransitionsMeta.json UpdateUserToGroup.json UpgradeResult.json User.json UserPickerResults.json UserWrite.json UsersAndGroups.json Version.json VersionIssueCounts.json VersionMove.json VersionUnresolvedIssueCounts.json Vote.json Watchers.json Workflow.json WorkflowMapping.json WorkflowScheme.json Worklog.json WorklogChangedSince.json WorklogIdsRequest.json WorklogWithPagination.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// CommentsWithPagination defined from schema:
// {
//   "title": "Comments With Pagination",
//   "id": "https://docs.atlassian.com/jira/REST/schema/comments-with-pagination#",
//   "type": "object",
//   "definitions": {
//     "user": {
//       "title": "User",
//       "type": "object",
//       "properties": {
//         "active": {
//           "type": "boolean"
//         },
//         "avatarUrls": {
//           "type": "object",
//           "patternProperties": {
//             ".+": {
//               "type": "string"
//             }
//           }
//         },
//         "displayName": {
//           "type": "string"
//         },
//         "emailAddress": {
//           "type": "string"
//         },
//         "key": {
//           "type": "string"
//         },
//         "name": {
//           "type": "string"
//         },
//         "self": {
//           "type": "string"
//         },
//         "timeZone": {
//           "type": "string"
//         }
//       }
//     }
//   },
//   "properties": {
//     "comments": {
//       "title": "comments",
//       "type": "array",
//       "items": {
//         "title": "Comment",
//         "type": "object",
//         "properties": {
//           "author": {
//             "title": "User",
//             "type": "object",
//             "properties": {
//               "active": {
//                 "type": "boolean"
//               },
//               "avatarUrls": {
//                 "type": "object",
//                 "patternProperties": {
//                   ".+": {
//                     "type": "string"
//                   }
//                 }
//               },
//               "displayName": {
//                 "type": "string"
//               },
//               "emailAddress": {
//                 "type": "string"
//               },
//               "key": {
//                 "type": "string"
//               },
//               "name": {
//                 "type": "string"
//               },
//               "self": {
//                 "type": "string"
//               },
//               "timeZone": {
//                 "type": "string"
//               }
//             }
//           },
//           "body": {
//             "title": "body",
//             "type": "string"
//           },
//           "created": {
//             "title": "created",
//             "type": "string"
//           },
//           "id": {
//             "title": "id",
//             "type": "string"
//           },
//           "properties": {
//             "title": "properties",
//             "type": "array",
//             "items": {
//               "title": "Entity Property",
//               "type": "object",
//               "properties": {
//                 "key": {
//                   "title": "key",
//                   "type": "string"
//                 },
//                 "value": {
//                   "title": "value"
//                 }
//               }
//             }
//           },
//           "renderedBody": {
//             "title": "renderedBody",
//             "type": "string"
//           },
//           "self": {
//             "title": "self",
//             "type": "string"
//           },
//           "updateAuthor": {
//             "title": "User",
//             "type": "object",
//             "properties": {
//               "active": {
//                 "type": "boolean"
//               },
//               "avatarUrls": {
//                 "type": "object",
//                 "patternProperties": {
//                   ".+": {
//                     "type": "string"
//                   }
//                 }
//               },
//               "displayName": {
//                 "type": "string"
//               },
//               "emailAddress": {
//                 "type": "string"
//               },
//               "key": {
//                 "type": "string"
//               },
//               "name": {
//                 "type": "string"
//               },
//               "self": {
//                 "type": "string"
//               },
//               "timeZone": {
//                 "type": "string"
//               }
//             }
//           },
//           "updated": {
//             "title": "updated",
//             "type": "string"
//           },
//           "visibility": {
//             "title": "Visibility",
//             "type": "object",
//             "properties": {
//               "type": {
//                 "title": "type",
//                 "type": "string"
//               },
//               "value": {
//                 "title": "value",
//                 "type": "string"
//               }
//             }
//           }
//         }
//       }
//     },
//     "maxResults": {
//       "title": "maxResults",
//       "type": "integer"
//     },
//     "startAt": {
//       "title": "startAt",
//       "type": "integer"
//     },
//     "total": {
//       "title": "total",
//       "type": "integer"
//     }
//   }
// }
type CommentsWithPagination struct {
	Comments   Comments `json:"comments,omitempty" yaml:"comments,omitempty"`
	MaxResults int      `json:"maxResults,omitempty" yaml:"maxResults,omitempty"`
	StartAt    int      `json:"startAt,omitempty" yaml:"startAt,omitempty"`
	Total      int      `json:"total,omitempty" yaml:"total,omitempty"`
}