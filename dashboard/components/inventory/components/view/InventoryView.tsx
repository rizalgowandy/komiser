import Image from 'next/image';
import { NextRouter } from 'next/router';
import formatNumber from '../../../../utils/formatNumber';
import providers, { Provider } from '../../../../utils/providerHelper';
import Button from '../../../button/Button';
import Checkbox from '../../../checkbox/Checkbox';
import AlertIcon from '../../../icons/AlertIcon';
import BookmarkIcon from '../../../icons/BookmarkIcon';
import Input from '../../../input/Input';
import Sidepanel from '../../../sidepanel/Sidepanel';
import SidepanelHeader from '../../../sidepanel/SidepanelHeader';
import SidepanelPage from '../../../sidepanel/SidepanelPage';
import SidepanelTabs from '../../../sidepanel/SidepanelTabs';
import { ToastProps } from '../../../toast/hooks/useToast';
import {
  HiddenResource,
  InventoryFilterData,
  InventoryStats,
  View
} from '../../hooks/useInventory/types/useInventoryTypes';
import InventoryFilterSummary from '../filter/InventoryFilterSummary';
import InventoryViewHeader from './InventoryViewHeader';
import InventoryViewAlerts from './alerts/InventoryViewAlerts';
import useViews from './hooks/useViews';

type InventoryViewProps = {
  filters: InventoryFilterData[] | undefined;
  displayedFilters: InventoryFilterData[] | undefined;
  setToast: (toast: ToastProps | undefined) => void;
  inventoryStats: InventoryStats | undefined;
  router: NextRouter;
  views: View[] | undefined;
  getViews: (edit?: boolean | undefined, viewName?: string | undefined) => void;
  hiddenResources: HiddenResource[] | undefined;
  setHideOrUnhideHasUpdate: (hideOrUnhideHasUpdate: boolean) => void;
};
function InventoryView({
  filters,
  displayedFilters,
  setToast,
  inventoryStats,
  router,
  views,
  getViews,
  hiddenResources,
  setHideOrUnhideHasUpdate
}: InventoryViewProps) {
  const {
    isOpen,
    openModal,
    closeModal,
    view,
    handleChange,
    saveView,
    loading,
    page,
    goTo,
    deleteView,
    bulkItems,
    bulkSelectCheckbox,
    onCheckboxChange,
    handleBulkSelection,
    unhideLoading,
    unhideResources,
    deleteLoading
  } = useViews({
    setToast,
    views,
    router,
    getViews,
    hiddenResources,
    setHideOrUnhideHasUpdate
  });

  return (
    <>
      <InventoryViewHeader
        openModal={openModal}
        views={views}
        router={router}
        saveView={saveView}
        setToast={setToast}
        loading={loading}
        deleteView={deleteView}
        deleteLoading={deleteLoading}
      />

      {/* Alerts button */}
      {router.query.view && (
        <div className="absolute right-0">
          <Button
            style="secondary"
            size="xs"
            transition={false}
            onClick={() => {
              openModal(undefined, 'alerts');
            }}
            loading={loading}
          >
            <AlertIcon width={20} height={20} />
            Alerts
          </Button>
        </div>
      )}

      {/* Save as a view button */}
      {!router.query.view && (
        <Button size="sm" onClick={() => openModal(filters)}>
          <BookmarkIcon width={20} height={20} />
          Save as a view
        </Button>
      )}

      {/* Sidepanel */}
      <Sidepanel isOpen={isOpen} closeModal={closeModal} noScroll={true}>
        <SidepanelHeader
          title={router.query.view ? view.name : 'Save as a view'}
          subtitle={`${inventoryStats?.resources} ${
            inventoryStats?.resources === 1 ? 'resource' : 'resources'
          } ${
            router.query.view
              ? 'are part of this view'
              : 'will be added to this view'
          }`}
          deleteAction={router.query.view ? () => deleteView(false) : undefined}
          deleteLabel="Delete view"
          closeModal={closeModal}
        />
        <SidepanelTabs
          goTo={goTo}
          page={page}
          tabs={
            router.query.view
              ? ['View', 'Alerts', 'Hidden Resources']
              : ['View']
          }
        />
        <SidepanelPage page={page} param="view">
          <form onSubmit={e => saveView(e)} className="flex flex-col gap-4">
            <div className="flex flex-col gap-2">
              {displayedFilters &&
                displayedFilters.length > 0 &&
                displayedFilters.map((data, idx) => (
                  <InventoryFilterSummary key={idx} data={data} />
                ))}
            </div>
            <Input
              name="name"
              label={router.query.view ? 'View name' : 'Choose a view name'}
              type="text"
              error="Please provide a name"
              value={view.name}
              action={handleChange}
              autofocus={true}
            />

            <div className="ml-auto">
              <Button
                size="lg"
                type="submit"
                loading={loading}
                disabled={!view.name}
              >
                {router.query.view ? 'Update view' : 'Save as a view'}{' '}
                <span className="flex items-center justify-center rounded-lg bg-black-900/20 py-1 px-2 text-xs">
                  {inventoryStats?.resources}
                </span>
              </Button>
            </div>
          </form>
        </SidepanelPage>

        <SidepanelPage page={page} param="alerts">
          <InventoryViewAlerts viewId={view.id} setToast={setToast} />
        </SidepanelPage>

        <SidepanelPage page={page} param="hidden resources">
          {hiddenResources && hiddenResources.length > 0 && (
            <>
              <div className="max-h-[calc(100vh-300px)] overflow-y-auto overflow-x-hidden">
                <table className="w-full table-auto bg-white text-left text-xs text-gray-900">
                  <thead className="bg-white">
                    <tr className="shadow-[inset_0_-1px_0_0_#cfd7d74d]">
                      <th className="py-4 px-2">
                        <div className="flex items-center">
                          <Checkbox
                            checked={bulkSelectCheckbox}
                            onChange={handleBulkSelection}
                          />
                        </div>
                      </th>
                      <th className="py-4 px-2">Cloud</th>
                      <th className="py-4 px-2">Service</th>
                      <th className="py-4 px-2">Name</th>
                      <th className="py-4 px-2">Region</th>
                      <th className="py-4 px-2">Account</th>
                      <th className="py-4 px-2 text-right">Cost</th>
                    </tr>
                  </thead>
                  <tbody>
                    {hiddenResources.map(item => (
                      <tr
                        key={item.id}
                        className={`border-b border-black-200/30 last:border-none ${
                          bulkItems &&
                          bulkItems.find(currentId => currentId === item.id)
                            ? 'border-black-200/70 bg-komiser-120'
                            : 'border-black-200/30 bg-white hover:bg-black-100/50'
                        } border-b last:border-none`}
                      >
                        <td className="py-4 px-2">
                          <Checkbox
                            checked={
                              bulkItems &&
                              !!bulkItems.find(
                                currentId => currentId === item.id
                              )
                            }
                            onChange={e => onCheckboxChange(e, item.id)}
                          />
                        </td>
                        <td className="py-4 pl-2 pr-6">
                          <div className="flex items-center gap-2">
                            <picture className="flex-shrink-0">
                              <img
                                src={providers.providerImg(
                                  item.provider as Provider
                                )}
                                className="h-6 w-6 rounded-full"
                                alt={item.provider}
                              />
                            </picture>
                            <span>{item.provider}</span>
                          </div>
                        </td>
                        <td className="py-4 px-2">{item.service}</td>
                        <td className="py-4 px-2">
                          <p className="... w-24 truncate">{item.name}</p>
                        </td>
                        <td className="py-4 px-2">
                          <p className="... w-24 truncate">{item.region}</p>
                        </td>
                        <td className="py-4 px-2">
                          <p className="... w-24 truncate">{item.account}</p>
                        </td>
                        <td className="py-4 px-2 text-right">
                          ${formatNumber(item.cost)}
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
              <div className="flex justify-end">
                <Button
                  size="lg"
                  disabled={bulkItems && bulkItems.length === 0}
                  loading={unhideLoading}
                  onClick={unhideResources}
                >
                  Unhide resources{' '}
                  <span className="flex items-center justify-center rounded-lg bg-white/10 py-1 px-2 text-xs">
                    {formatNumber(bulkItems.length)}
                  </span>
                </Button>
              </div>
            </>
          )}

          {hiddenResources && hiddenResources.length === 0 && (
            <div className="rounded-lg bg-black-100 p-6">
              <div className="flex flex-col items-center gap-6">
                <Image
                  src="/assets/img/purplin/dashboard.svg"
                  alt="Purplin"
                  width={150}
                  height={100}
                />
                <div className="flex flex-col items-center justify-center gap-2 px-24 text-center">
                  <p className="font-semibold text-black-900">
                    No hidden resources in this view
                  </p>
                  <p className="text-sm text-black-400">
                    To hide a resource from this view, select and hide them on
                    the inventory table.
                  </p>
                </div>
              </div>
            </div>
          )}
        </SidepanelPage>
      </Sidepanel>
    </>
  );
}

export default InventoryView;
