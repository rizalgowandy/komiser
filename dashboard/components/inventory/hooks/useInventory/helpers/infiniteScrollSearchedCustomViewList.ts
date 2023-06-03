import { NextRouter } from 'next/router';
import { SetStateAction } from 'react';
import settingsService from '../../../../../services/settingsService';
import { ToastProps } from '../../../../toast/hooks/useToast';
import { InventoryItem, View } from '../types/useInventoryTypes';

type InfiniteScrollSearchedCustomViewListProps = {
  router: NextRouter;
  shouldFetchMore: boolean;
  isVisible: boolean;
  views: View[] | undefined;
  query: string;
  batchSize: number;
  skippedSearch: number;
  setToast: (value: SetStateAction<ToastProps | undefined>) => void;
  setSearchedInventory: (
    value: SetStateAction<InventoryItem[] | undefined>
  ) => void;
  setShouldFetchMore: (value: SetStateAction<boolean>) => void;
  setSkippedSearch: (value: SetStateAction<number>) => void;
};

/** Load the next 50 results when the user scrolls a searched custom view list to the end */
function infiniteScrollSearchedCustomViewList({
  router,
  shouldFetchMore,
  isVisible,
  views,
  query,
  batchSize,
  skippedSearch,
  setToast,
  setSearchedInventory,
  setShouldFetchMore,
  setSkippedSearch
}: InfiniteScrollSearchedCustomViewListProps) {
  if (
    shouldFetchMore &&
    isVisible &&
    query &&
    router.query.view &&
    views &&
    views.length > 0
  ) {
    const id = router.query.view;
    const filterFound = views.find(view => view.id.toString() === id);

    if (filterFound) {
      const payloadJson = JSON.stringify(filterFound?.filters);

      settingsService
        .getInventory(
          `?limit=${batchSize}&skip=${skippedSearch}&query=${query}&view=${id}`,
          payloadJson
        )
        .then(res => {
          if (res.error) {
            setToast({
              hasError: true,
              title: `There was an error fetching more resources!`,
              message: `Please refresh the page and try again.`
            });
          } else {
            setSearchedInventory(prev => {
              if (prev) {
                return [...prev, ...res];
              }
              return res;
            });
            setSkippedSearch(prev => prev + batchSize);

            if (res.length >= batchSize) {
              setShouldFetchMore(true);
            } else {
              setShouldFetchMore(false);
            }
          }
        });
    }
  }
}

export default infiniteScrollSearchedCustomViewList;
