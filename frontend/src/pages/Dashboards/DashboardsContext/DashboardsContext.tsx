import { FetchResult } from '@apollo/client'
import { UpsertDashboardMutation } from '@graph/operations'
import * as Types from '@graph/schemas'
import {
	DashboardDefinition,
	DashboardMetricConfigInput,
	Maybe,
} from '@graph/schemas'
import { createContext } from '@util/context/context'
import { useState } from 'react'

interface DashboardsContext {
	allAdmins: Maybe<
		{ __typename?: 'Admin' } & Pick<
			Types.Admin,
			'id' | 'name' | 'email' | 'photo_url'
		>
	>[]
	dashboards: Maybe<DashboardDefinition>[]
	updateDashboard: ({
		id,
		name,
		metrics,
		layout,
	}: {
		id?: string
		name: string
		metrics: DashboardMetricConfigInput[]
		layout?: string
	}) => Promise<FetchResult<UpsertDashboardMutation>>
}

export const [useDashboardsContext, DashboardsContextProvider] =
	createContext<DashboardsContext>('Dashboards')

const DashboardsContextProvider = ({ children }) => {
	const [dashboards, setDashboards] = useState<Maybe<DashboardDefinition>[]>([])

	const updateDashboard = async ({
		id,
		name,
		metrics,
		layout,
	}: {
		id?: string
		name: string
		metrics: DashboardMetricConfigInput[]
		layout?: string
	}) => {
		const result = await upsertDashboardMutation({
			variables: { id, name, metrics, layout },
		})

		const newDashboard = result.data?.upsertDashboard
		if (newDashboard) {
			const dashboardMap = new Map(dashboards.map(d => [d.id, d]));
			dashboardMap.set(newDashboard.id, newDashboard);
			const uniqueDashboards = Array.from(dashboardMap.values());
			setDashboards([...uniqueDashboards, newDashboard])
		}

		return result
	}

	return (
		<DashboardsContext.Provider value={{ dashboards, updateDashboard }}>
			{children}
		</DashboardsContext.Provider>
	)
}

export { useDashboardsContext, DashboardsContextProvider }
