import { useEffect, useRef, useState } from 'react';
import Head from 'next/head';
import { useRouter } from 'next/router';

import CloudAccountItem from '@components/cloud-account/components/CloudAccountItem';
import Toast from '@components/toast/Toast';
import Modal from '@components/modal/Modal';
import CloudAccountsHeader from '@components/cloud-account/components/CloudAccountsHeader';
import CloudAccountsLayout from '@components/cloud-account/components/CloudAccountsLayout';

import useCloudAccount from '@components/cloud-account/hooks/useCloudAccounts/useCloudAccount';
import CloudAccountsSidePanel from '@components/cloud-account/components/CloudAccountsSidePanel';
import CloudAccountDeleteContents from '@components/cloud-account/components/CloudAccountDeleteContents';
import useToast from '@components/toast/hooks/useToast';

function CloudAccounts() {
  const [editCloudAccount, setEditCloudAccount] = useState<boolean>(false);
  const [isDeleteModalOpen, setIsDeleteModalOpen] = useState<boolean>(false);

  const { toast, setToast, dismissToast } = useToast();
  const router = useRouter();

  const currentViewProvider = router.query.view as string;

  const {
    cloudAccounts,
    setCloudAccounts,
    openModal,
    cloudAccountItem,
    setCloudAccountItem,
    page,
    goTo,
    isNotCustomView,
    isLoading
  } = useCloudAccount();

  const [filteredCloudAccounts, setFilteredCloudAccounts] =
    useState(cloudAccounts);

  useEffect(() => {
    if (!currentViewProvider) setFilteredCloudAccounts(cloudAccounts);
    else {
      setFilteredCloudAccounts(
        cloudAccounts.filter(
          account =>
            account.provider.toLowerCase() === currentViewProvider.toLowerCase()
        )
      );
    }
  }, [currentViewProvider, cloudAccounts]);

  const closeRemoveModal = () => {
    setIsDeleteModalOpen(false);
  };

  if (!cloudAccounts || isLoading) return null;

  return (
    <>
      <Head>
        <title>Cloud Accounts - Komiser</title>
        <meta name="description" content="Cloud Accounts - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      {/* Wraps the cloud account page and handles the custom views sidebar */}
      <CloudAccountsLayout router={router} cloudAccounts={cloudAccounts}>
        <CloudAccountsHeader isNotCustomView={isNotCustomView} />

        {filteredCloudAccounts.map(account => (
          <CloudAccountItem
            key={account.id}
            account={account}
            openModal={openModal}
            setCloudAccountItem={setCloudAccountItem}
            setIsDeleteModalOpen={setIsDeleteModalOpen}
            setEditCloudAccount={setEditCloudAccount}
          />
        ))}
      </CloudAccountsLayout>

      {/* Delete Modal */}
      <Modal isOpen={isDeleteModalOpen} closeModal={() => closeRemoveModal()}>
        <div className="flex max-w-xl flex-col gap-y-6 p-8 text-black-400">
          {cloudAccountItem && (
            <CloudAccountDeleteContents
              cloudAccount={cloudAccountItem}
              onCancel={closeRemoveModal}
              setToast={setToast}
            />
          )}
        </div>
      </Modal>

      {cloudAccountItem && (
        <CloudAccountsSidePanel
          isOpen={editCloudAccount}
          closeModal={() => setEditCloudAccount(false)}
          cloudAccount={cloudAccountItem}
          cloudAccounts={cloudAccounts}
          setCloudAccounts={setCloudAccounts}
          setToast={setToast}
          page={page}
          goTo={goTo}
        />
      )}

      {/* Toast component */}
      {toast && <Toast {...toast} dismissToast={dismissToast} />}
    </>
  );
}

export default CloudAccounts;
