/* post-processed by ./scripts/process-types.py */
/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface DocumentOut {
  id: string;
  path: string;
  title: string;
}

export interface Group {
  createdAt: Date;
  currency: string;
  id: string;
  name: string;
  updatedAt: Date;
}

export interface GroupUpdate {
  currency: string;
  name: string;
}

export interface ItemAttachment {
  createdAt: Date;
  document: DocumentOut;
  id: string;
  type: string;
  updatedAt: Date;
}

export interface ItemAttachmentUpdate {
  title: string;
  type: string;
}

export interface ItemCreate {
  description: string;
  labelIds: string[];

  /** Edges */
  locationId: string;
  name: string;
}

export interface ItemField {
  booleanValue: boolean;
  id: string;
  name: string;
  numberValue: number;
  textValue: string;
  timeValue: string;
  type: string;
}

export interface ItemOut {
  attachments: ItemAttachment[];
  createdAt: Date;
  description: string;

  /** Future */
  fields: ItemField[];
  id: string;
  insured: boolean;
  labels: LabelSummary[];

  /** Warranty */
  lifetimeWarranty: boolean;

  /** Edges */
  location: LocationSummary;
  manufacturer: string;
  modelNumber: string;
  name: string;

  /** Extras */
  notes: string;
  purchaseFrom: string;

  /** @example 0 */
  purchasePrice: string;

  /** Purchase */
  purchaseTime: Date;
  quantity: number;
  serialNumber: string;
  soldNotes: string;

  /** @example 0 */
  soldPrice: string;

  /** Sold */
  soldTime: Date;
  soldTo: string;
  updatedAt: Date;
  warrantyDetails: string;
  warrantyExpires: Date;
}

export interface ItemSummary {
  createdAt: Date;
  description: string;
  id: string;
  insured: boolean;
  labels: LabelSummary[];

  /** Edges */
  location: LocationSummary;
  name: string;
  quantity: number;
  updatedAt: Date;
}

export interface ItemUpdate {
  description: string;
  fields: ItemField[];
  id: string;
  insured: boolean;
  labelIds: string[];

  /** Warranty */
  lifetimeWarranty: boolean;

  /** Edges */
  locationId: string;
  manufacturer: string;
  modelNumber: string;
  name: string;

  /** Extras */
  notes: string;
  purchaseFrom: string;

  /** @example 0 */
  purchasePrice: string;

  /** Purchase */
  purchaseTime: Date;
  quantity: number;

  /** Identifications */
  serialNumber: string;
  soldNotes: string;

  /** @example 0 */
  soldPrice: string;

  /** Sold */
  soldTime: Date;
  soldTo: string;
  warrantyDetails: string;
  warrantyExpires: Date;
}

export interface LabelCreate {
  color: string;
  description: string;
  name: string;
}

export interface LabelOut {
  createdAt: Date;
  description: string;
  id: string;
  items: ItemSummary[];
  name: string;
  updatedAt: Date;
}

export interface LabelSummary {
  createdAt: Date;
  description: string;
  id: string;
  name: string;
  updatedAt: Date;
}

export interface LocationCreate {
  description: string;
  name: string;
}

export interface LocationOut {
  createdAt: Date;
  description: string;
  id: string;
  items: ItemSummary[];
  name: string;
  updatedAt: Date;
}

export interface LocationOutCount {
  createdAt: Date;
  description: string;
  id: string;
  itemCount: number;
  name: string;
  updatedAt: Date;
}

export interface LocationSummary {
  createdAt: Date;
  description: string;
  id: string;
  name: string;
  updatedAt: Date;
}

export interface PaginationResultRepoItemSummary {
  items: ItemSummary[];
  page: number;
  pageSize: number;
  total: number;
}

export interface UserOut {
  email: string;
  groupId: string;
  groupName: string;
  id: string;
  isOwner: boolean;
  isSuperuser: boolean;
  name: string;
}

export interface UserUpdate {
  email: string;
  name: string;
}

export interface ServerResult {
  details: any;
  error: boolean;
  item: any;
  message: string;
}

export interface ServerResults {
  items: any;
}

export interface ServerValidationError {
  field: string;
  reason: string;
}

export interface UserRegistration {
  email: string;
  name: string;
  password: string;
  token: string;
}

export interface ApiSummary {
  build: Build;
  demo: boolean;
  health: boolean;
  message: string;
  title: string;
  versions: string[];
}

export interface Build {
  buildTime: string;
  commit: string;
  version: string;
}

export interface ChangePassword {
  current: string;
  new: string;
}

export interface GroupInvitation {
  expiresAt: Date;
  token: string;
  uses: number;
}

export interface GroupInvitationCreate {
  expiresAt: Date;
  uses: number;
}

export interface ItemAttachmentToken {
  token: string;
}

export interface TokenResponse {
  expiresAt: Date;
  token: string;
}
