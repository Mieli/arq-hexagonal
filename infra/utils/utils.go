package utils

import (
	"time"

	pkgutils "delegacia.com.br/app/domain/utils"
)

func SetEventRecord(eventRecord *pkgutils.EventRecord, option string, active bool) *pkgutils.EventRecord {
	today := time.Now()
	if eventRecord != nil {
		if option == pkgutils.CREATED_AT {
			eventRecord.CreatedAt = &today
		}
		if option == pkgutils.UPDATED_AT {
			eventRecord.UpdatedAt = &today
		}
		if option == pkgutils.DELETED_AT {
			eventRecord.DeletedAt = &today
		}
		eventRecord.Active = active
	}

	return eventRecord
}

func MergeEventRecord(newEventRecord, oldEventRecord pkgutils.EventRecord) pkgutils.EventRecord {
	newEventRecord.Active = oldEventRecord.Active
	newEventRecord.CreatedAt = oldEventRecord.CreatedAt
	newEventRecord.UpdatedAt = oldEventRecord.UpdatedAt
	newEventRecord.DeletedAt = oldEventRecord.DeletedAt
	return newEventRecord

}
