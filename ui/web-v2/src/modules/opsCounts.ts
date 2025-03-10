import {
  createSlice,
  createEntityAdapter,
  createAsyncThunk
} from '@reduxjs/toolkit';

import * as autoOpsGrpc from '../grpc/autoops';
import { OpsCount } from '../proto/autoops/ops_count_pb';
import {
  ListOpsCountsRequest,
  ListOpsCountsResponse
} from '../proto/autoops/service_pb';

import { AppState } from '.';

const MODULE_NAME = 'opsCounts';

export const opsCountsAdapter = createEntityAdapter({
  selectId: (opsCount: OpsCount.AsObject) => opsCount.id
});

export const { selectAll, selectById } = opsCountsAdapter.getSelectors();

export interface ListOpsCountsParams {
  environmentId: string;
  ids: Array<string>;
}

export const listOpsCounts = createAsyncThunk<
  ListOpsCountsResponse.AsObject,
  ListOpsCountsParams | undefined,
  { state: AppState }
>(`${MODULE_NAME}/listOpsCounts`, async (params) => {
  const request = new ListOpsCountsRequest();
  request.setAutoOpsRuleIdsList(params.ids);
  request.setEnvironmentId(params.environmentId);
  const result = await autoOpsGrpc.listOpsCounts(request);
  return result.response.toObject();
});

const initialState = opsCountsAdapter.getInitialState<{
  loading: boolean;
}>({
  loading: false
});

export type OpsCountsState = typeof initialState;

export const opsCountsSlice = createSlice({
  name: MODULE_NAME,
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(listOpsCounts.pending, (state) => {
        state.loading = true;
      })
      .addCase(listOpsCounts.fulfilled, (state, action) => {
        opsCountsAdapter.removeAll(state);
        opsCountsAdapter.upsertMany(state, action.payload.opsCountsList);
        state.loading = false;
      });
  }
});
