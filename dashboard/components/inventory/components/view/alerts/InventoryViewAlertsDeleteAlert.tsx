import Image from 'next/image';
import Button from '../../../../button/Button';
import { AlertMethod, Alert } from './hooks/useAlerts';
import useEditAlerts from './hooks/useEditAlerts';
import { ToastProps } from '../../../../toast/hooks/useToast';

type InventoryViewAlertsDeleteAlertProps = {
  alertMethod: AlertMethod;
  closeAlert: (action?: 'hasChanges' | undefined) => void;
  viewId: number;
  setToast: (toast: ToastProps | undefined) => void;
  viewControllerOnCancelButton: () => void;
  currentAlert: Alert | undefined;
};

function InventoryViewAlertsDeleteAlert({
  alertMethod,
  viewId,
  closeAlert,
  setToast,
  viewControllerOnCancelButton,
  currentAlert
}: InventoryViewAlertsDeleteAlertProps) {
  const { deleteAlert, loading } = useEditAlerts({
    alertMethod,
    currentAlert,
    viewId,
    closeAlert,
    setToast
  });
  return (
    <div className="rounded-lg bg-komiser-100 p-6">
      <div className="flex flex-col items-center gap-6">
        <Image
          src="/assets/img/others/warning.svg"
          alt="Purplin"
          width={48}
          height={48}
          className="mx-auto flex-shrink-0"
        />
        <div className="mb-8 flex flex-col items-center gap-2 px-4">
          <p className="text-center font-semibold text-black-900">
            Are you sure you want to delete this alert?
          </p>
          <p className="text-center text-sm text-black-400">
            By deleting the “{currentAlert?.name}”{' '}
            {currentAlert?.isSlack ? 'slack' : 'webhook'} alert, you won’t
            receive any more notifications regarding the cost limit you set up.
          </p>
        </div>
      </div>
      <div className="flex items-center justify-end">
        <div className="flex gap-4">
          <Button
            style="ghost"
            size="lg"
            onClick={viewControllerOnCancelButton}
          >
            Cancel
          </Button>
          <Button
            size="sm"
            style="delete"
            type="button"
            onClick={() => {
              viewControllerOnCancelButton();
              if (currentAlert) {
                deleteAlert(currentAlert.id);
              }
            }}
            loading={loading}
          >
            Delete alert
          </Button>
        </div>
      </div>
    </div>
  );
}

export default InventoryViewAlertsDeleteAlert;
